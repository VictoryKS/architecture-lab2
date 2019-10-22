// This file contains examples of scenarios implementation using
// the SDK for balancers management.

const balancers = require('./balancers/client');

const client = balancers.Client('http://localhost:8080');

// Scenario 1: Display selected balancer.
client.getBalancer(1)
    .then((res) => {
        console.log('=== Scenario 1 ===');
        console.log('Balancer1:');
        console.log(res);
    })
    .catch((e) => {
        console.log(`Problem listing balancer's state: ${e.message}`);
    });

// Scenario 2: Change machine's state.
client.changeState(2, 1)
    .then((resp) => {
        console.log('=== Scenario 2 ===');
        console.log('Update machine\'s state:', resp);
        console.log('Success!');
    })
    .catch((e) => {
        console.log(`Problem changing machine's state: ${e.message}`);
    });
