#!/usr/bin/env node

var amqp = require('amqplib/callback_api');
const updataServer = require("./server")
let handelMsg = require('./handle')

const opt = {credentials: require('amqplib').credentials.plain('shop_1', '123456')};

const amqpConnect = async ()=>{
    await updataServer.updataServer()
    amqp.connect('amqp://106.52.26.252/%2Fshop_1?heartbeat=3', opt, function (error0, connection) {
        if (error0) {
            throw new Error('lalal');
        }
        // console.log('connection', connection);
        connection.createChannel(function (error1, channel) {
            if (error1) {
                throw new Error('lalal');
                console.log(" error 2");
            }
            var exchange = 'siyu_order_exchange';
            channel.assertExchange(exchange, 'direct', {
                durable: true
            });
            var queue = 'debug_siyu_order';
            channel.assertQueue(queue, {
                durable: true
            });
            channel.bindQueue(queue, exchange, 'debug_siyu_order');
            console.log(" [*] Waiting for messages in %s. To exit press CTRL+C", queue);
            channel.consume(queue,async function (msg) {
                console.log(" [x] Received %s", msg.content.toString());
                
                let re = await handelMsg.handleAmqp(msg.content)
                if(re) {
                    //确认ACK
                    console.log(" 确认ACK ");
                    channel.ack(msg);
                }else{
                    console.log(" 打印出错 ");
                }
                // console.log("测试环境全部确认ack");
                // channel.ack(msg);
            }, {
                noAck: false
            });
        });
    });
}

try {
    amqpConnect()
} catch (error) {
    console.log(error)
}


console.log('end')
