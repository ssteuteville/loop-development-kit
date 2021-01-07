using System;
using Grpc.Core;
using Grpc.Core.Interceptors;
using NUnit.Framework;
using OliveHelpsLDK.Filesystem;
using OliveHelpsLDK.Logging;

namespace OliveHelpsLDK.Test.FileSystem.Integration
{
    [TestFixture]
    public class FileSystemTests
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
            _client = new FilesystemClient(channel, session, _logger);
        }

        private IFilesystemService _client;
        private readonly ILogger _logger = new Logger("integration-test-logger");

        [Test]
        public void GetsLowLevelErrorForQueries()
        {
            var task = _client.QueryDirectory("/tmp");

            Assert.Throws<AggregateException>(() => task.Wait());
        }

        [Test]
        public void GetsLowLevelErrorForStreams()
        {
            var call = _client.StreamDirectory("/tmp");

            var nextStatus = call.MoveNext();

            Assert.Throws<AggregateException>(() => nextStatus.Wait());
        }
    }
}