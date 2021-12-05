const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        listForums: () => client.get("/"),
        registerUser: (name, interests) => client.post("/", { name, interests }),
    }
};

module.exports = { Client };
