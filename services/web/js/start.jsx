import React from "react";
import ReactDOM from "react-dom/client";
import Cookies from 'universal-cookie';
var TicTacToe = require('tic_tac_toe');

var api = new TicTacToe.DefaultApi(new TicTacToe.ApiClient("http://localhost:8081"));

class Game extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            currentPlayer: '',
            winner: {},
            grid: []
        };
        this.loadState();       
    }

    loadState = () => {
        let g = this;
        api.currentPlayer(this.props.id, function (error, data, response) {
            if (error) {
                console.error(error);
            } else {
                g.setState({currentPlayer: data});
            }
        })
        api.winner(this.props.id, function (error, data, response) {
            if (error) {
                console.error(error);
            } else {
                console.log(data);
                g.setState({winner: data});
            }
        });
        api.gameGrid(this.props.id,
            function (error, data, response) {
                if (error) {
                    console.error(error);
                } else {
                    g.setState({grid: data.grid});
                }
            });
      };

    render() {
        let finished = this.state.winner.winner || this.state.winner.draw;
        return <>
            <GameHeader
            winner={this.state.winner}
            currentPlayer={this.state.currentPlayer}
            />
            <Grid 
            id={this.props.id}
            grid={this.state.grid}
            play={(x,y) => {
                if (!finished) {
                    const KeyPlayerTokenX = "TicTacToePlayerToken-X";
                    const KeyPlayerTokenO = "TicTacToePlayerToken-O";
                    var keyPlayerToken = KeyPlayerTokenO
                    if (this.state.currentPlayer == 'X') {
                        keyPlayerToken = KeyPlayerTokenX;
                    }
                    const cookies = new Cookies();
                    var playerToken = cookies.get(keyPlayerToken);
                    console.log(playerToken);
                    var playerTokenBuf = atob(playerToken);
                    console.log(playerTokenBuf);
                    console.log(x);
                    console.log(y);
                    var g = this;
                    api.play(this.props.id, playerTokenBuf, x, y, function(error, data, response){
                        if (error) {
                            console.error(error);
                        } else {
                            console.log(data);
                            g.loadState();
                        }
                    });
                }
            }}
            />
        </>
    }
}

function GameHeader(props) {
    var text = "";
    if (props.winner.winner) {
        text = "Winner: " + props.winner.winner;
    } else if (props.winner.draw) {
        text = "Draw!";
    } else {
        text = "Next Player: " + props.currentPlayer;
    }
    return <div className="grid grid-cols-2 my-5">
    <a className="hover:underline justify-self-start text-blue-800 block" href="/">Home</a>

    <div className="block justify-self-end">
        {text}
    </div>
</div>
}

class Grid extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        return <div className="grid grid-cols-3 gap-2 my-5">
                <GridInner 
                grid={this.props.grid}
                play={this.props.play}/>
            </div>;
    }
}

function GridInner(props) {
    if (props.grid) {
        var spaces = [];
        props.grid.forEach((row, i) => {
            row.forEach((space, j) => {
                spaces.push(
                    <GridSpace
                        key={"space:" + i + "," + j}
                        content={space}
                        play={() => {
                           props.play(i,j);
                        }}
                    />);
            });
        });
        return <>{spaces}</>;
    } else {
        return <p>No grid</p>;
    }
}

function GridSpace(props) {
    var content = <>&nbsp;&nbsp;</>;
    if (props.content != "") {
        content = props.content;
    }
    return <div className="block">
        <button
            className="block w-full rounded-md text-white font-extrabold text-center hover:bg-blue-600 bg-blue-500 p-7 hover:shadow-lg focus:outline-none"
            type="button"
            onClick={props.play}
        >
            {content}
        </button>
    </div>;
}

var grid = document.getElementById("game");
if (grid) {
    var gid = grid.dataset.gameid;
    ReactDOM.createRoot(grid).render(<Game id={gid}/>);
}