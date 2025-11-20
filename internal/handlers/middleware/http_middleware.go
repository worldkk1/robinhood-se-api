package middleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	redisrate "github.com/go-redis/redis_rate/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type Middleware func(http.Handler) http.HandlerFunc

type AuthUser struct {
	UserId string
	Role   string
}

type RedisRateLimiter struct {
	*redisrate.Limiter
}

type contextKey string

const ContextUserKey contextKey = "user"
const rateLimitRequestKey = "rate_limit_request"

func MiddlewareChain(middleware ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middleware) - 1; i >= 0; i-- {
			next = middleware[i](next)
		}

		return next.ServeHTTP
	}
}

func LoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Println(time.Since(start), r.Method, r.URL.Path)
	}
}

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenSplit := strings.Split(authHeader, "Bearer ")
		if len(tokenSplit) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := tokenSplit[1]
		secretKey := os.Getenv("SECRET_KEY")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			return []byte(secretKey), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		sub, ok := claims["sub"].(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		role, ok := claims["role"].(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		authUser := AuthUser{
			UserId: sub,
			Role:   role,
		}
		ctx := context.WithValue(r.Context(), ContextUserKey, authUser)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func CheckRoleAdminMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(ContextUserKey).(AuthUser)
		if !ok {
			http.Error(w, "User not found", http.StatusInternalServerError)
			return
		}
		// Check if role_id is admin role
		if user.Role != "ae4c58a6-101a-4b0b-a63e-e187d1920c7e" {
			http.Error(w, "no_permission", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func SetupRedisRateLimiter() *redisrate.Limiter {
	connString := os.Getenv("REDIS_CONNECTION_STRING")
	client := redis.NewClient(&redis.Options{
		Addr: connString,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	return redisrate.NewLimiter(client)
}

func RateLimitMiddleware(next http.Handler) http.HandlerFunc {
	redisRateLimiter := SetupRedisRateLimiter()
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := redisRateLimiter.Allow(r.Context(), rateLimitRequestKey, redisrate.Limit{
			Rate:   1,
			Burst:  10,
			Period: time.Second,
		})
		if res.Allowed <= 0 {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	}
}
