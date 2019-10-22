const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        getBalancer: (id) => client.get('/balancers?id=' + id),
        changeState: (id, isWorking) => client.post('/balancers', {  id,  isWorking })
    }

};

module.exports = { Client };
