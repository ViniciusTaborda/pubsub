## PubSub

This repository contains a simple implementation of the Publish-Subscribe pattern in Go.

## Overview

The Publish-Subscribe pattern is a messaging pattern where publishers send messages to topics, and subscribers receive messages from topics they have subscribed to. This implementation provides a simple in-memory implementation that allows publishers to send messages to topics and subscribers to receive messages from topics.

## Possible Improvements
 - Error handling in message publish
 - Message publishing retry
 - Make the logging in file system more flexible and customizable

## Code Coverage
To view the code coverage, follow these steps:

 - Run the following command at the root folder of this project:
 ```
 make run tests
 ```
- This will run the tests and export the coverage information to a `coverage.html` file. 
- Open it in any web browser
- See example:

![image](https://github.com/ViniciusTaborda/pubsub/assets/72284753/5dbff54c-3306-49cd-8fc1-737b3e42dbbf)


## License

This project is licensed under the MIT License. See the LICENSE file for more information.
