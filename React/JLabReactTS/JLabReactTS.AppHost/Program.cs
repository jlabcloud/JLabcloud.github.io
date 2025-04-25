var builder = DistributedApplication.CreateBuilder(args);

builder.AddProject<Projects.JLabReactTS_Server>("jlabreactts-server");

builder.Build().Run();
