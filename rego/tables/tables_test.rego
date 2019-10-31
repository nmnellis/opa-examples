package tables

users = {
  "nick": {
    "tables": [
      {
        "name": "users",
        "operations": [
          "read",
        ]
      },
      {
        "name": "permissions",
        "operations": [
          "read",
        ]
      }
    ]
  },
  "torrin": {
    "tables": [
      {
        "name": "users",
        "operations": [
          "read",
          "update",
          "delete"
        ],
        "rowIds": [
          0,
          1,
        ]
      }
    ]
  }
}

test_nick_read_user {
  allow with input as {"table": "users", "operation": "read", "rowId":1234, "user": "nick"} with data.users as users
}
test_nick_create_user {
    not allow with input as {"table": "users", "operation": "create", "rowId":1234, "user": "nick"} with data.users as users
}
test_nick_read_permission {
    allow with input as {"table": "permissions", "operation": "read", "rowId":0, "user": "nick"} with data.users as users
}
test_nick_create_permission {
    not allow with input as {"table": "permissions", "operation": "create", "rowId":0, "user": "nick"} with data.users as users
}

test_torrin_read_user_1 {
    allow with input as {"table": "users", "operation": "read", "rowId": 1, "user": "torrin"} with data.users as users
}
test_torrin_update_user_1 {
    allow with input as {"table": "users", "operation": "update", "rowId": 1, "user": "torrin"} with data.users as users
}
test_torrin_read_user_10 {
    not allow with input as {"table": "users", "operation": "read", "rowId": 10, "user": "torrin"} with data.users as users
}
test_torrin_update_user_10 {
    not allow with input as {"table": "users", "operation": "update", "rowId": 10, "user": "torrin"} with data.users as users
}
test_torrin_create_user {
    not allow with input as {"table": "users", "operation": "create", "rowId":0, "user": "torrin"} with data.users as users
}

test_torrin_read_permission {
    not allow with input as {"table": "permissions", "operation": "read", "rowId":0, "user": "torrin"} with data.users as users
}
test_torrin_create_permission {
    not allow with input as {"table": "permissions", "operation": "create", "rowId":0, "user": "torrin"} with data.users as users
}