FROM registry.access.redhat.com/ubi9
WORKDIR /opt
RUN dnf update -y
RUN dnf install wget -y
RUN dnf install dotnet-sdk-7.0 dotnet-runtime-7.0 -y
RUN dotnet new console -o network --use-program-main
WORKDIR /opt/network
COPY Program.cs /opt/network/Program.cs
RUN dotnet publish --runtime linux-x64 -p:PublishSingleFile=true --self-contained=true
