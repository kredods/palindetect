# Palindetect
Feed application that detects when messages are palindromes

Instance running on http://13.57.28.144:3000/ 

# Setup
Clone the repository
```git clone https://github.com/kredods/palindetect.git```


# Run the application
```cd palindetect```
``` docker-compose -f docker-compose.yml up --build```
The application should be running on port 3000 of your container's IP


# API Docs
Creating a new message
![Create Message](https://github.com/kredods/palindetect/raw/master/PalinDetect%20(Create%20message).jpg)

- URL: `/v1/messages`

- Method: `POST`

- Data:
```
{
	"body": "kayak"
}
```

- Success :

  - Code: `201`

  - Success Response:

```
{
    "body": "kayak",
    "isPalindrome": true,
    "id": "a1b2c3b4c5d5e77"
}
```

Fetch all messages
![Get Messages](https://github.com/kredods/palindetect/raw/master/PalinDetect%20(GET%20ALL).jpg)
- URL: `/v1/messages`

- Method: `GET`

- Success :

  - Code: `200`

  - Success Response:

```
[
  {
      "body": "kayak",
      "isPalindrome": true,
      "id": "a1b2c3b4c5d5e77"
  },
  {
      "body": "kayaka",
      "isPalindrome": false,
      "id": "a1b2c3b4c5d5e78"
  },
  {
      "body": "mom",
      "isPalindrome": true,
      "id": "a1b2c3b4c5d5e79"
  }
]
```
Deleting a message
- URL: `/v1/messages/{id}`

- Method: `DELETE`

- Success :

  - Code: `200`

Updating a message
![Update Message](https://github.com/kredods/palindetect/raw/master/PalinDetect%20(Update%20message).jpg)
- URL: `/v1/messages`

- Method: `PUT`

- Data:
```
{
	"body": "kayak"
}
```

- Success :

  - Code: `200`
  
  
 Get by Id
![Get Message by Id](https://github.com/kredods/palindetect/raw/master/PalinDetect%20(GET%20by%20id).jpg)

- URL: `/v1/messages/{id}`

- Method: `GET`


- Success :

  - Code: `200`

  - Success Response:

```
{
    "body": "kayak",
    "isPalindrome": true,
    "id": "a1b2c3b4c5d5e77"
}
```
  
