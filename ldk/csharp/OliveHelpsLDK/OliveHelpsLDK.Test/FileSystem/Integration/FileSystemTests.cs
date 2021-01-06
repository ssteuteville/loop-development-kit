using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Grpc.Core;
using Grpc.Core.Interceptors;
using NUnit.Framework;
using OliveHelpsLDK.Filesystem;
using OliveHelpsLDK.Logging;
using Proto;
using FileInfo = OliveHelpsLDK.Filesystem.FileInfo;

namespace OliveHelpsLDK.Test.FileSystem.Integration
{
    [TestFixture]
    public class FileSystemTests
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
            Client = new FilesystemClient(channel, session, Logger);
        }

        private IFilesystemService Client;
        private readonly ILogger Logger = new Logger("integration-test-logger");
        
        [Test]
        public void GetsLowLevelErrorForQueries()
        {
            var task = Client.QueryDirectory("/tmp");

            Assert.Throws<AggregateException>(() => task.Wait());
        }
        
        [Test]
        public void GetsLowLevelErrorForStreams()
        {
            var call = Client.StreamDirectory("/tmp");

            var nextStatus = call.MoveNext();

            Assert.Throws<AggregateException>(() => nextStatus.Wait());
        }
    }
    
}