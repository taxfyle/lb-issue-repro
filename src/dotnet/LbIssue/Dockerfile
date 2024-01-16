FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS base
USER $APP_UID
WORKDIR /app
EXPOSE 5001
EXPOSE 5002

FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
ARG BUILD_CONFIGURATION=Release
WORKDIR /src
COPY ["LbIssue.Grpc/LbIssue.Grpc.csproj", "LbIssue.Grpc/"]
COPY ["LbIssue.ApiHost/LbIssue.ApiHost.csproj", "LbIssue.ApiHost/"]
RUN dotnet restore "LbIssue.ApiHost/LbIssue.ApiHost.csproj"
COPY . .
WORKDIR "/src/LbIssue.ApiHost"
RUN dotnet build "LbIssue.ApiHost.csproj" -c $BUILD_CONFIGURATION -o /app/build

FROM build AS publish
ARG BUILD_CONFIGURATION=Release
RUN dotnet publish "LbIssue.ApiHost.csproj" -c $BUILD_CONFIGURATION -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "LbIssue.ApiHost.dll"]
