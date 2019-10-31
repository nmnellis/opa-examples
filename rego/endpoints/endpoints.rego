package endpoints

default allow = false

user = user {
  user := data.users[input.user]
}

# match method and regex of path
allow {
  some i
  re_match(user.endpoints[i].regex,input.path) # match regex to input path
  user.endpoints[i].methods[_] == input.method # user has matching method
}