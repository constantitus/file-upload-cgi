## Input sample


From environment variables, we are using `CONTENT_TYPE` to get our boundary (separator)
`CONTENT_TYPE="multipart/form-data; boundary=-----------------------------364047128122568367761297282173"`

From stdin, we get the username, password and file name + content (times the number of files) separated by the boundary that we got earlier
```
-----------------------------364047128122568367761297282173
Content-Disposition: form-data; name="user"

MyUser
-----------------------------364047128122568367761297282173
Content-Disposition: form-data; name="pass"

MyPassword
-----------------------------364047128122568367761297282173
Content-Disposition: form-data; name="the_file"; filename="test.txt"
Content-Type: text/plain

TESTFILE #1
THIS IS THE first TEXT FILE


-----------------------------364047128122568367761297282173
Content-Disposition: form-data; name="the_file"; filename="test2.txt"
Content-Type: text/plain

TESTFILE #2
THIS IS THE second TEXT FILE


-----------------------------364047128122568367761297282173--
 
```
