// This file contains examples of scenarios implementation using
// the SDK for channels management.

const channels = require("./channels/client");

const client = channels.Client("http://localhost:8080");

// Scenario 1: Display available forums.
client
  .listForums()
  .then((list) => {
    console.log(`=== Scenario 1 === \nForums: \n${list} `);
  })
  .catch((e) => {
    console.log(`There is a problem listing forums: ${e.message}`);
  });

// Scenario 2: Register new user.
client
  .registerUser("testUser", ["films"])
  .then((resp) => {
    console.log(`=== Scenario 2 === \nUser response: ${resp}`);
  })
  .catch((e) => {
    console.log(`There is a problem registering new user: ${e.message}`);
  });