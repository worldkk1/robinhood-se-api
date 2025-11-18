# API for manage card and comment

## Feature
- Create card
- Get card list
- Get card detail
- Edit card detail
- Add comment for any card
- Edit comment
- Delete comment
- Get card history
- Update card status
- Archive card
- Authenticate and authorize user

## Model
```
tasks
  id
  title
  description
  status
  user_id
  archive_at
  created_at
  updated_at

task_comments
  id
  comment
  user_id
  created_at
  updated_at

task_logs
  id
  task_id
  title
  description
  status
  created_at
  updated_at

users
  id
  name
  email
  role_id
  created_at
  updated_at

roles
  id
  name
  created_at
  updated_at
```
