## DBS


A simple key-value in memory database written in golang using custom protocol for client/server communication. 

The motivation behind this project is learning how to make custom protocols and golang.

### Protocol Commands

| ID         | Sent by | Description                           |
|------------|---------|---------------------------------------|
| `CONN`     | Client  | Establish a connection with a database |
| `CREATE`   | Client  | Create an entry in database           |
| `READ`     | Client  | Read an entry from database           |
| `UPDATE`   | Client  | Update an entry in database           |
| `DELETE`   | Client  | Delete an entry from database         |
| `OK`       | Server  | Command acknowledgement               |
| `ERR`      | Server  | Error                                 |
| `CLOSECONN`| Client  | Close a connection                    |

---

**CONN**

A client can connect to the server using the `CONN` command. It takes the name of the database as an required argument.

Syntax:

`CONN <db_name>`

If the database doesn't exist, one is simply created.

---

**CREATE**

Create an entry(key/value) pair to the database. It takes the key name, follwed by the value length and the value to be stored.

`CREATE <key-name> <length>\r\n[value]`

where:
* `key-name`: Is the key 
* `length`: Length of value
* `value`: Value to be stored in db

---

**READ**

Read a value from the database. The server will reply with the value associated with the key.

`READ <key-name>`

---

**UPDATE**

Update a value of a key in database. The server will reply with `OK` if value is successfully updated or `ERR` with error message if no key with given name is found or any error in server.

`UPDATE <key-name> <length>\r\n[value]`

where:
* `key-name`: Name of key
* `length`: Length of value
* `value`: Value to be updated

---

**DELETE**

Delete a key from the database. The server will reply with `OK` if value is successfully deleted or if the key doesn't exists.

`DELETE <key-name>`

---

**OK/ERR**

When a server recieves a command it can either reply with OK or ERR

`OK` doesn't have any text after it.

`ERR <error-message>` if the format of the errors returned by the server.
A server will not close connection in case of `ERR`

---

**CLOSECONN**

Close a connection with exisitng database and server. 

Syntax:

`CLOSECONN`

---


### RESOURCES:

1. [Custom protocol design](https://ieftimov.com/posts/understanding-bytes-golang-build-tcp-protocol/)
