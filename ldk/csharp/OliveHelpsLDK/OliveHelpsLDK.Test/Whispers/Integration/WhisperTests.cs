using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Grpc.Core;
using Grpc.Core.Interceptors;
using NUnit.Framework;
using OliveHelpsLDK.Filesystem;
using OliveHelpsLDK.Logging;
using OliveHelpsLDK.Whispers;
using Proto;
using FileInfo = OliveHelpsLDK.Filesystem.FileInfo;

namespace OliveHelpsLDK.Test.FileSystem.Integration
{
    [TestFixture]
    public class WhisperTests
    {
        [OneTimeSetUp]
        public void ConnectClient()
        {
            var channel = new Channel("localhost:4770", ChannelCredentials.Insecure)
                .Intercept(new ExceptionLoggingInterceptor(Logger));
            
            var session = new Session
            {
                LoopId = "LOOP",
                Token = "TOKEN"
            };
            Client = new WhisperClient(channel, session, Logger);
        }

        private IWhisperService Client;
        private readonly ILogger Logger = new Logger("integration-test-logger");
        
        [Test]
        public void GetsLowLevelErrorForWrites()
        {
            var task = Client.MarkdownAsync(new WhisperMarkdown()
            {
                Config = new WhisperConfig() { Label = "hello" },
                Markdown = "#HEADER"
            });

            Assert.Throws<AggregateException>(() => task.Wait());
        }
    }
    
}