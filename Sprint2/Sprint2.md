**List of unit tests for frontend**
- Default path set to home test case: When all asynchronous tasks are completed, the test checks to see if the default path that our web application routes to is to the home page. It does this with the expect() and toBe() functions. 
- Home button test case: All buttons in the webpage are found with the queryAll() function. The buttons are then stored in an array. Our test then retrieves the first button in our array, which is the home button. The test clicks on the button and checks to see if it routes to the home page. 
- Officers button test case: All buttons in the webpage are found with the queryAll() function. The buttons are then stored in an array. Our test then retrieves the second button in our array, which is the officers button. The test clicks on the button and checks to see if it routes to the officers page. 
- Rent a pc button test case: All buttons in the webpage are found with the queryAll() function. The buttons are then stored in an array. Our test then retrieves the third button in our array, which is the rent a pc button. The test clicks on the button and checks to see if it routes to the rent a pc page. 
- Login button test case: All buttons in the webpage are found with the queryAll() function. The buttons are then stored in an array. Our test then retrieves the fourth button in our array, which is the login button. The test clicks on the button and checks to see if it routes to the login page. 
- Sign up button test case: All buttons in the webpage are found with the queryAll() function. The buttons are then stored in an array. Our test then retrieves the fifth button in our array, which is the sign up button. The test clicks on the button and checks to see if it routes to the sign up page. 

**List of unit tests for Cypress**
- Visit localhost test: This test is an end-to-end test that simply checks if localhost:4200 is accessible. Localhost:4200 must be running via the command ‘ng serve’ for this test to be successful. The home page for the Society of PC Building web application is displayed on the Cypress testing window.
- Find text in the SPCB web app test: This is an end-to-end test that checks whether the correct text displays on the SPCB web application. After the browser loads, the test will attempt to find text that contains ‘HOME’. If the text is not found, the test fails, but if the text is found, the interactive element that contains the text ‘HOME’ (in our case it is the button) is clicked. The test repeats this process for the rest of the buttons on the navbar. After reaching the sign up page, the test will check to see if the current page has the text ‘First Name’. It continues to check for ‘Last Name’, ‘Email’, and ‘Password’ until the test completes. 
- Check club name text test: This is a component test that checks whether our club’s name ‘The Society of PC Building’ is displayed on our navbar. The test finds the correct data-cy attribute labeled ‘title’ and checks to see if it contains our club name.
- Home button test: This is a component test that checks whether the home is clicked. The test finds the correct data-cy attribute labeled ‘home’ and clicks it. The test then checks to see if the change property contains the appropriate string that was returned by clicking the home button. 

**Other work completed for the frontend**
- Created a carousel component that contains an array of pictures from our club’s GBMs. After 8 seconds, the carousel switches to a new picture. This was done with the autoSlideImages() function which clicks the nextClick() function after the default interval passes. The user can also view a different picture by clicking the left arrow, right arrow, or indicator at the bottom. This was done with the prevClick(), nextClick(), and selectImage() functions that would switch to a new index when it is clicked. 
- Updated the officers page to have some information about current board members. Information listed on the officers page can include their position, picture, name, year, and major. Each officer’s information is put on a mat-card element, which are UI components imported from Angular Materials.
- Changed color scheme of the web page from purple with yellow highlights to black and blue highlights.



**BACKEND DOCUMENTATION**

Packages:
- main - Thusfar, the only 'real' package. Contains all of the API endpoints and handles all of the routing. This will need to be better organized in the future; we need to place the various functions in other packages based on their functionality.

Structs:
- User - A gorm model struct which is used in the SQLite database to represent a user and their credentials. The primary key has been manually changed to be the user's email address.
- Env - A struct with a single field: a pointer to a gorm.DB instance. This exists so that we don't have to manually open the database in every function that needs it.

 Functions:
- getHash - A very simple function that passes in a []byte and returns the corresponding hash generated by bcrypt. Called by userLogin and userRegister.GenerateJWT - Returns a JSON web token (JWT) signed using HS256. Tokens are set to expire 24 hours after generation. Called by userLogin. test - Prints "Test success" when called. Called by the /api/test endpoint.
- userRegister - An Env struct behavior. Creates a new entry in the User table of the database provided that no other entry with the same email exists. Called by /api/signup endpoint.
- userLogin - An Env struct behavior. Compares hashes of the provided password string with the corresponding entry in the User table provided such an entry exists. If the hashes match, a JWT is returned along with a success message. Called by /api/login endpoint.
- deactivateUser - An Env struct behavior. Deletes the user with provided email address from the User table provided they exist. This deletion is a hard delete using a raw SQL statement, hence the deletion is permanent. Called by /api/deactivate-account endpoint.
- updateUser - An Env struct behavior. Updates the user's account information with the provided changes provided that the user exists in the User table. Does not update the email address. Called by /api/update-account endpoint.
 
 **List of unit tests for backend**
- TestUserLogin: The unit test case tests the UserLogin() function in an Env struct using a mock environment that includes a SQLite database and a mock User       object. It creates an HTTP POST request to the login endpoint with the mock User object as the body and an HTTP response recorder to capture the response. Finally, it checks the status code, content type, and response body of the HTTP response to ensure that the function returns the expected result.
- TestUserRegister:This is a unit test that checks the userRegister() function's functionality in an Env struct. It creates a mock environment, including a SQLite database, and defines a mock User object with valid credentials. The test does this by sending an HTTP POST request to the register endpoint.
- TestDeactivateUser: This unit test case tests the functionality of deactivating a user in the application. The first test case checks that an existing user is deleted from the database, and the second test case checks that an error message is returned when attempting to delete a non-existent user.
- TestUpdateUser:in this unit test there are two test cases are performed: 1) updating an existing user should update the user in the database, and 2) updating a non-existent user should return an error message. The function creates a mock environment with an in-memory SQLite database, and uses the "Env" struct to call the "UpdateUser" function. The tests check that the user information is updated correctly and that the expected error messages are returned.
      
**Middleware Handlers**
ValidateJWT - Checks that the value mapped to the "Token" key in a passed in http request is a valid JWT. Writes a response message telling the caller that they are unauthorized if the token is invalid or if there is no token at all. If the token is valid, the passed in function is allowed to execute. Endpoints which are wrapped by this handler: /api/test and /api/deactivate-account.
    API Endpoints:
      /api/test - Calls the test function. Uses the ValidateJWT handler. Accepts http GET requests.
      /api/signup - Calls the userRegister function. No additional middleware handlers. Accepts http POST requests.
      /api/login - Calls the userLogin function. No additional middleware handlers. Accepts http POST requests.
      /api/deactivate-account - Calls the deactivateUser function. Uses the ValidateJWT handler. Accepts http DELETE requests.
      /api/update-account - Calls the updateUser function. Uses the Validate JWT handler. Accepts http PUT requests.
  main_test - Contains all of the unit tests written for the backend.
Database - SPCB.db:
    Tables:
        Users - A table containing user account details. Implements gorm.model. Primary key is set to email address. Additional fields include: firstname, lastname, and password.
        
Link to narrated video presentation: https://www.youtube.com/watch?v=kpe-e2aXhbw
