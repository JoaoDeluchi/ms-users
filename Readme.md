# Prerequisites

Docker installed (https://www.docker.com/products/docker-desktop/)

# Building the Docker Image

Open a terminal and navigate to your project directory.

# Build the Docker image using the following command:

``` Bash
docker build -t zpe-user-management-api .
```

# Running the API in a Docker Container

Start the API in a detached mode using the following command:

```  Bash
docker run -d -p 8080:8080 zpe-user-management-api
```
-d runs the container in detached mode, allowing the terminal to prompt for further commands.
-p 8080:8080 maps port 8080 on the host machine to port 8080 inside the container. This allows you to access the API running in the container from your browser at http://localhost:8080. You can adjust the port numbers if needed.


# Using the API with Postman

To test the API, you can use the provided Postman collection. Follow the steps below:

## Import the Postman Collection

1. Open Postman.
2. Click on the `Import` button in the top left corner.
3. Select `File` and upload the `ZPE_User_Management_API.postman_collection.json` file included in the project repository.

## Using the Collection

1. After importing, you will see a new collection named `ZPE User Management API` in your Postman sidebar.
2. Expand the collection to see the available endpoints:
    - `Create User`
    - `Get All Users`
    - `Get User by ID`
    - `Update User Roles`
    - `Delete User`

## Example Requests

### Create User
- Method: `POST`
- URL: `http://localhost:8080/users`
- Body (raw JSON):
```json
{
"Name": "Lebron James",
"Role": ["Admin"],
"Email": "papai_lebron@lakers.com"
}
```

### Get All Users
- Method: `Get`
- URL: `http://localhost:8080/users`


### Get User by id
- Method: `Get`
- URL: `http://localhost:8080/${id}`
```note
The id is auto generated and you can get using the Get All route
```

### Update User
- Method: `Put`
- URL: `http://localhost:8080/users/${id}`
- Body (raw JSON):
```json
{
    "roles": [
        "Modifier"
    ]
}
```
```note
The id is auto generated and you can get using the Get All route
```


### Delete User
- Method: `Put`
- URL: `http://localhost:8080/users/${id}`
```note
The id is auto generated and you can get using the Get All route
```
