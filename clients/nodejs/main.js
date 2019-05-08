var PROTO_PATH = __dirname + '/../../pb/users/users.proto';
var grpc = require('grpc');
var protoLoader = require('@grpc/proto-loader');

// Suggested options for similarity to existing grpc.load behavior
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH, {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
var protoDescriptor = grpc.loadPackageDefinition(packageDefinition);

// The protoDescriptor object has the full package hierarchy
var users = protoDescriptor.users;

var stub = new users.UserInfoService('localhost:2333', grpc.credentials.createInsecure());
var resp = stub.GetUserInfo({
        name: "wmh"
    },

    function (err, userInfo) {
        if (err) {
            console.log(err);
        } else {
            console.log(userInfo);
        }
    });