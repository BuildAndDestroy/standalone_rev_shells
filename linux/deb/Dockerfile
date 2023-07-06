FROM debian:latest
WORKDIR /opt
RUN apt update -y
RUN apt install wget -y
RUN wget https://packages.microsoft.com/config/debian/11/packages-microsoft-prod.deb -O packages-microsoft-prod.deb
RUN apt install ./packages-microsoft-prod.deb -y
RUN rm packages-microsoft-prod.deb
RUN apt update -y
RUN apt install dotnet-sdk-7.0 dotnet-runtime-7.0 -y
RUN dotnet new console -o network --use-program-main
WORKDIR /opt/network
COPY Program.cs /opt/network/Program.cs
RUN dotnet publish --runtime linux-x64 -p:PublishSingleFile=true --self-contained=true
