using System;
using RabbitMQ.Client;
using System.Text;

class Send
{
    public static void Main()
    {
        var factory = new ConnectionFactory() { HostName = "192.168.31.71" };
        using (var connection = factory.CreateConnection())
        using (var channel = connection.CreateModel())
        {
            channel.QueueDeclare(queue: "hello",
                                 durable: false,
                                 exclusive: false,
                                 autoDelete: false,
                                 arguments: null);

            string message = "Hello World!";
            var body = Encoding.UTF8.GetBytes(message);

            channel.BasicPublish(exchange: "",
                                 routingKey: "hello",
                                 basicProperties: null,
                                 body: body);
            Console.WriteLine(" [x] Sent {0}", message);
            Console.WriteLine("请输入内容:");
            string str = Console.ReadLine();
            while (str != "quit" || str != "exit")
            {
                channel.BasicPublish(exchange: "",
                        routingKey: "hello",
                        basicProperties: null,
                        body: Encoding.UTF8.GetBytes(str));
                str =DateTime.Now.ToString();
                Thread.Sleep(1000);
            }
        }

        Console.WriteLine(" Press [enter] to exit.");
        Console.ReadLine();
    }
}