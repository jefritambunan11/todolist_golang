Backend Service Aplikasi Todo List, Sebagai Portfolio Sederhana

Author : Jefri Tambunan

###################
User API - Register
###################
Method: POST

localhost:8080/api/users

JSON Body : 
{
    "name": "Jefri",    
    "email": "jefri@gmail.com",
    "password": "password"
}

################
User API - Login
################
Method: POST

localhost:8080/api/sessions

JSON Body : 
{
  "email": "jefri@gmail.com",
  "password": "password" 
}

################################################################
User API - Check Email Sudah Digunakan Atau Belum
################################################################
Method: POST

localhost:8080/api/email_checkers

JSON Body : 
{
    "email": "jefri@gmail.com"
}





####################
Todo API - List All
####################
Method: GET

localhost:8080/api/todo_list


########################################
Todo API - List All With Page Navigation
########################################
Method: GET

localhost:8080/api/todo_list?page=1




##############################
Todo API - List Detail Pake ID
##############################
Method: GET

localhost:8080/api/todo_list/1


#################
Todo API - Create
#################

Method: POST

localhost:8080/api/todo

set Bearer Token - Token Didapat Ketika Login Berhasil
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJKZWZyaSIsInVzZXJfcGFzc3dvcmQiOiIkMmEkMDQkWUIzOWtUZ2VFUnFEcXpNWEs0YVpLZWhKVnlEemF1QXBzdHF5L1hQajl0dzBPeDcuWDVPVEMifQ.dDiBcMl3iaAQJ9jBpJPvZagmAFlEb9UAdAyBJtpfqXM


JSON Body : 
{
    "todo": "Menonton TV",
    "date_time": "2022-07-08T19:00:00+07:00" 
}


#################
Todo API - Update
#################

Method: PUT

localhost:8080/api/todo/1

set Bearer Token - Token Didapat Ketika Login Berhasil
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJKZWZyaSIsInVzZXJfcGFzc3dvcmQiOiIkMmEkMDQkWUIzOWtUZ2VFUnFEcXpNWEs0YVpLZWhKVnlEemF1QXBzdHF5L1hQajl0dzBPeDcuWDVPVEMifQ.dDiBcMl3iaAQJ9jBpJPvZagmAFlEb9UAdAyBJtpfqXM


JSON Body : 
{
    "todo": "Menonton TV Satelit",
    "date_time": "2022-07-08T23:00:00+07:00"    
}



#################
Todo API - Delete
#################

Method: DELETE

localhost:8080/api/todo/1

set Bearer Token - Token Didapat Ketika Login Berhasil
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiJKZWZyaSIsInVzZXJfcGFzc3dvcmQiOiIkMmEkMDQkWUIzOWtUZ2VFUnFEcXpNWEs0YVpLZWhKVnlEemF1QXBzdHF5L1hQajl0dzBPeDcuWDVPVEMifQ.dDiBcMl3iaAQJ9jBpJPvZagmAFlEb9UAdAyBJtpfqXM

