version_settings(
    check_updates=True,
    constraint='>=0.22.2'
) 
load('ext://namespace', 'namespace_yaml')

update_settings(max_parallel_updates=5)

# Command-line flags
config.define_bool("bare",args=False,usage="If set, runs resources as bare processes without kubernetes")
config.define_bool("bots",args=False,usage="If set, run a set of bots to automatically play games")

args = config.parse()
bare = "bare" in args and args["bare"]
bots = "bots" in args and args["bots"]

local_resource(
    'tests',
    cmd='go test --short ./...',
    deps=['.'],
    ignore=['.output'],
    labels=["test"],
    allow_parallel=True,
)

local_resource(
    'tests (long)',
    cmd='go test ./...',
    trigger_mode = TRIGGER_MODE_MANUAL,
    auto_init = False,
    labels=["test"],
    allow_parallel=True,
)

local_resource(
    'play',
    cmd='go run ./services/bot/cmd/onegame http://localhost:8081',
    trigger_mode = TRIGGER_MODE_MANUAL,
    auto_init = False,
    labels=["test"]
)

if bare:
    local_resource(
        "mongodb",
        serve_cmd="docker run --rm -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=password -p 27017:27017 mongo:4.0.8",
    )
else:
    lightstep_access_token=""

    if os.path.exists("secrets.yaml"):
        secrets = read_yaml("secrets.yaml")
        lightstep_access_token=secrets["lightstep"]["access_token"]

    k8s_yaml(namespace_yaml('tictactoe'))

    # Load the base Helm chart for all resources
    k8s_yaml(helm(
        'charts/tic-tac-toe',
        namespace='tictactoe',
        set=[
            "mongodb.statefulset=true",
            ],
    ))

    # Load the Tilt support Helm chart
    k8s_yaml(helm(
        'tilt/charts/otel',
        namespace='tictactoe',
        set=[
            "lightstep.access_token=" + lightstep_access_token,
            ],
    ))
    k8s_resource("jaeger", port_forwards="16686:16686")

    k8s_resource("mongodb-standalone", port_forwards="27017:27017")

    local_resource("base-image", "make dockerbaseimage")

def server(name, port_forwards=[], port="8080", grpcui_port="8081"):
    if bare:
        local_resource(
            name,
            cmd='make ' + name + "_local",
            serve_cmd="cd .build/" + name  + "/ && ./app_local",
            serve_env={
                "PORT": str(port),
                "GRPCUI_PORT": str(grpcui_port),
                "MONGO_CONN": "mongodb://admin:password@localhost:27017",
            },
            deps = ["services/" + name, "common"],
            labels=[name]
        )
    else:
        local_resource(
            name+"-build",
            'make ' + name,
            deps = ["services/" + name, "common"],
            labels=[name]
        )
        docker_build(
            "docker.io/tictactoverengineered/"+name,
            '.build/' + name  + "/",
            dockerfile = "docker/app/Dockerfile",
            live_update = [
                sync('.build/' + name + '/app', '/root/app'),
                run('./restart.sh'),
            ]
        )

        k8s_resource(
            name, 
            port_forwards=port_forwards, 
            resource_deps=[name+'-build', 'base-image'],
            labels=[name]
        )

def space(name,ports=[]):
    if bare:
        local_resource(
            name + "-build",
            cmd='make ' + name + "_local",
            deps = ["services/" + name, "common"],
            labels=["space"]
        )
        for i in range(0, 3):
            for j in range(0, 3):
                local_resource(
                    name + " (" + str(i) + "," + str(j) + ")",
                    serve_cmd="cd .build/" + name  + "/ && ./app_local",
                    serve_env={
                        "PORT": "80" + str(i) + str(j),
                        "XPOS": str(i),
                        "YPOS": str(j),
                        "MONGO_CONN": "mongodb://admin:password@localhost:27017",
                    },
                    deps = ["services/" + name, "common"],
                    resource_deps=[name + "-build"],
                    labels=["space"]
                )
    else:
        local_resource(
            name+"-build",
            'make ' + name,
            deps = ["services/" + name, "common"],
            labels=["space"]
        )
        docker_build(
            "docker.io/tictactoverengineered/"+name,
            '.build/' + name  + "/",
            dockerfile = "docker/app/Dockerfile",
            live_update = [
                sync('.build/' + name + '/app', '/root/app'),
                run('./restart.sh'),
            ]
        )

        k8s_resource(name, resource_deps=[name+'-build'], labels=["space"])

server("api", port_forwards="8081:8080",port=8081)
server("web", port_forwards="8080:8080",port=8080)
server("gamerepo", port_forwards=["8082:8080", "8083:8081"], port=8082,grpcui_port=8083)
server("currentturn", port_forwards=["8084:8080", "8085:8081"], port=8084,grpcui_port=8085)
server("grid", port_forwards=["8086:8080", "8087:8081"], port=8086,grpcui_port=8087)
server("checker", port_forwards=["8088:8080", "8089:8081"], port=8088,grpcui_port=8089)
server("turncontroller", port_forwards=["8090:8080", "8091:8081"], port=8090,grpcui_port=8091)
server("matchmaker", port_forwards=["8092:8080", "8093:8081"], port=8092,grpcui_port=8093)
space("space")

# Bots
if bots or config.tilt_subcommand == "down":
    if not bare:
        # Load the Bots Helm chart
        k8s_yaml(helm(
            'charts/tic-tac-toe-bots',
            namespace='tictactoe',
        ))

    server("bot",port_forwards=["2112:2112"], port=2112)
