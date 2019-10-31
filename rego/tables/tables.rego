package tables

default allow = false

user = user {
  user := data.users[input.user]
}

# you have general access to the database table
allow {
  some i
  not user.tables[i].rowIds # no rowIds provided on the table
  input.table == user.tables[i].name # table name matches
  input.operation == user.tables[i].operations[_] # user has matching operation

}

# you only have access to specific rows in the table
allow {
  some i
  input.rowId == user.tables[i].rowIds[_] # rowId is compared to incomming rowId
  input.table == user.tables[i].name # table name matches
  input.operation == user.tables[i].operations[_] # user has matching operation
}
