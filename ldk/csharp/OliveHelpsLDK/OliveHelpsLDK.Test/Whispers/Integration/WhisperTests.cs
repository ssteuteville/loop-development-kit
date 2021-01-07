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

namespace OliveHelpsLDK.Test.Whispers.Integration
{
    [TestFixture]
    public class WhisperTests
    {
        [OneTimeSetUp]
        public void ConnectClient()
        {
            var channel = new Channel("localhost:4770", ChannelCredentials.Insecure)
                .Intercept(new ExceptionLoggingInterceptor(_logger));
            
            var session = new Session
            {
                LoopId = "LOOP",
                Token = "TOKEN"
            };
            _client = new WhisperClient(channel, session, _logger);
        }

        private IWhisperService _client;
        private readonly ILogger _logger = new Logger("integration-test-logger");
        
        [Test]
        public void GetsLowLevelErrorForWrites()
        {
            var task = _client.MarkdownAsync(new WhisperMarkdown()
            {
                Config = new WhisperConfig() { Label = "hello" },
                Markdown = "#HEADER"
            });

            Assert.Throws<AggregateException>(() => task.Wait());
        }
    }
    
}