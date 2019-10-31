package endpoints

users = {
  "nick": {
    "endpoints": [
      {
        "regex": ".*",
        "methods": [
          "GET"
        ]
      }
    ]
  },
  "torrin": {
    "endpoints": [
      {
        "regex": "^\/users\/([0-9]{1,})$",
        "methods": [
          "GET",
          "POST",
          "DELETE",
          "PUT"
        ]
      }
    ]
  }
}

test_nick_GET_anywhere_1 {
    allow with input as {"path": "/some/path/to/nowhere", "method": "GET", "user": "nick"} with data.users as users
}
test_nick_GET_anywhere_2 {
    allow with input as {"path": "/some/path/to/nowhere", "method": "GET", "user": "nick"} with data.users as users
}
test_nick_POST_anywhere_1 {
    not allow with input as {"path": "/some/path/to/nowhere", "method": "POST", "user": "nick"} with data.users as users
}
test_nick_POST_anywhere_2 {
    not allow with input as {"path": "/some/path/to/nowhere", "method": "POST", "user": "nick" } with data.users as users
}

test_torrin_GET_user {
    allow with input as {"path": "/users/1234", "method": "GET", "user": "torrin"} with data.users as users
}
test_torrin_GET_user_permissions {
    not allow with input as {"path": "/users/1234/permissions", "method": "GET", "user": "torrin"} with data.users as users
}
test_torrin_POST_user {
    allow with input as {"path": "/users/1234", "method": "POST", "user": "torrin"} with data.users as users
}
test_torrin_POST_user_permissions {
    not allow with input as {"path": "/users/1234/permissions", "method": "POST", "user": "torrin" } with data.users as users
}