package playroom

//Game interface for game logic shared by al game types
type Game interface {
	Prepare(t *Table, gametype string) //t is subtype of game (texas holdem or ... ), func prepares decka and initial table stuff, maybe also broadcast pool stuff?
	Start()                            //initial broadcast and deal
	Pause(s int)                       //pause actions after all open/waiting actions for s seconds
	CurrentState()                     //get current state Wait, Pause, Action, Finished
	Results(t *Table)                  //get results after finishing
}
