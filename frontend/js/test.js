const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const insecureCreds = grpc.credentials.createInsecure();

const PROTO_PATH = '/Users/random_number/yt-dl-rpc/api/rpc/rpc.proto';
const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });

const rpcProto = grpc.loadPackageDefinition(packageDefinition).rpc;
const client = new rpcProto.RPCService('localhost:3033', insecureCreds);

const defaultName = "https://www.youtube.com/watch?v=Q2YKXiTRcBE";

const call = client.Exec({
    URL: defaultName, 
    path: "/downloads", 
    Rename: "hashtag_from_js.webm", 
}, (err, response) => {
    if (err) {
        console.error('Error occurred: ', err);
    } else {
        console.log('Greeting: ', response);
    }
});