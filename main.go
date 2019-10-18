package main

import (
	"time"
	"fmt"
     	"strconv"
        "strings"
	"log"
        "os"
       	"net/http"
        "github.com/gorilla/mux"
        "database/sql"
      _ "github.com/lib/pq"
       
)

//var addr = flag.String("addr", ":8080", "http service address")

// Create room manager
type Room struct {
        Id   string `json:"Id"`
	Name  string `json:"Name"`
        CreateAt      int64  `json:"CreateAt"`
        Teamid        string `json:"Teamid"`
	Type          string `json:"type"`
	Displayname   string `json:"Displayname"`
        *Author
        *Session
        Authorid      string `json:"authorid"`
        //connections    map[*Connection]bool
}

type roomReq struct {
	// Name of the lobby the request goes to.
	name string

                
        room *Room
	// Reference to the connection, which requested.
	conn *Connection
}

type managedRoom struct {
	// Reference to room.
	room *Room

        
	// Member-count to allow removing of empty lobbies.
	count uint
}

type raamReq struct {
	// Name of the lobby the request goes to.
	namee string

                
        room Room
	// Reference to the connection, which requested.
	conn Connection
}

type Channel []Room

var appTeam *Team 

type Emoji struct {
        Id         string `json:"id"` 
	Keyi       string `json:"keyi"`
	Valeur     string `json:"valeur"`
	Descriptor string `json:"descriptor"`
}


type Post struct {
	Id            string          `json:"id"`
	Createat      int64           `json:"createat"`
	Userid        string          `json:"userid"`
	Channelid     string          `json:"channelid"`
	Rootid        string          `json:"rootid"`
	Message       string          `json:"message"`
	Color         string          `json:"color"`
        Typo          string          `json:"typo"`
        *Props        
}

type Team struct {
	Id              string `json:"Id"`
	Name            string `json:"Name"`
	Description     string `json:"Description"`
	Email           string `json:"Email"`
        Displayname     string `json:"Displayname"`
	CreateAt        int64  `json:"CreateAt"`
	namee           string `json:"namee"`
}

type Session struct {
	ID     int64   `json:"id"`
	Authori string `json:"authori"`
	Hash   string  `json:"hash"`
        
}

type TeamMember struct {
	Teamid   string `json:"team_id"`
	Userid   string `json:"user_id"`
	Roles    string `json:"roles"`
	DeleteAt int64  `json:"delete_at"`
}


type managedTeam struct {
	// Reference to room.
        
	teams map[string][]*Team
        rooms map[string][]*Room
	// Member-count to allow removing of empty lobbies.
	
}

type managedTeamCount struct {
	// Reference to room.
        count int
	teams map[string][]*Team
        rooms map[string][]*Room
	// Member-count to allow removing of empty lobbies.
	
}


type User struct {
    Id             int 
    Email          string
    CreatedAt      time.Time
    UpdatedAt     time.Time 
    HashedPassword []byte
    Password       string
}

// Every join and leave request will specify the name of the room
// the connection wants to join or leave.
func escaped(b byte) int {
	return strings.IndexByte("\\!\"#$%&'()*+,./:;<=>?@[]^_`{|}~-", b) 
}

func testEsc(s string) string {
     bi := []byte(s) 
     y := make([]byte, 0)
     
     for _, b := range bi {
           if escaped(b) > -1 {
              
             //return bi
           } else {
             y = append(y, b)
             //return y
           }
     }
     return string(y) 

}

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

// Emit the msg event to every member of the To-Room with the provided message content.
func (h *Hub) LoginQuat() func(http.ResponseWriter, *http.Request){
       
	return func(w http.ResponseWriter, r *http.Request) {
            vars := mux.Vars(r)  
            //_, _, lir, _ := store.GetTeamsAlli()
             fmt.Println("i am THE room", vars["room"])
              ua := r.Header.Get("User-Agent")
             fmt.Println("i am THE agent", ua)           

paz, _ := store.GetRoomsById()
log.Println("custom chat", paz)

_, _, top, zas := store.GetAlliByArray(paz)
log.Println("custom chat", zas)
poo, oor := store.GetTeams()
pao, oar := store.GetRooms()
log.Println("rooom", oor)
log.Println("rooom", oar)        
log.Println("rooom", poo)
log.Println("rooom", pao)

//h.GetTeamis(pao, poo)
//h.Teamis = top




h.Teamis = top
log.Println("custom chat", h.Teamis)
           for _, v := range h.Teamis {
              vars := mux.Vars(r) 
              fmt.Println("Hello", v)  
              for j, k := range v { 
              fmt.Println("Hello", k.teams) 
                       for _, kui := range k.teams {
                            fmt.Println("Hello",kui[0].Name)
                            
                            for _,pio := range kui {
                                fmt.Println("teamtyuuuuuuu", pio)
                                team := pio
                                if vars["team"] == team.Displayname {
                                         appTeam = team  
                                         fmt.Println("im a the team", vars["team"])  
                                         ///http.ServeFile(w, r, "assets/chat.html")
                                  } else {
                                       //w.WriteHeader(http.StatusNotFound)
                                        log.Println("custom 404")

                                  } 

                            }




                       }  
              fmt.Println("Hello", k.rooms) 
              //fmt.Println("Hellonnnnnnnnnn", k.rooms[i.Name]) 
              //fmt.Println("Hello", k.rooms[i.Name][0]) 
              //fmt.Println("Hello", k.rooms[i.Name][1])  
              fmt.Println("Hello", j)   
                       for o, ku := range k.rooms {
                            vars := mux.Vars(r) 
                            fmt.Println("team2222",o)  /// team
                            for _, kug := range ku {
                                fmt.Println("Hellonnnnnnbbbbbbbb",kug)
                                room := kug
                                
                                //fmt.Println("i am THE room", vars["room"])
                              if appTeam.Displayname == vars["team"] && vars["room"] == room.Displayname && appTeam.Id == room.Teamid {
                                    fmt.Println(vars["room"])
                                    fmt.Println("teamtyuuuuuuuggggggggggggggggggggggggggg", appTeam) 
                                    http.ServeFile(w, r, "assets/chat.html")
                              } else {
                                       ///w.WriteHeader(http.StatusNotFound)
                                        log.Println("custom 404nnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn")

                              } 

                            }
                       }          

              }

        }

    

            
     }

}



// If a room is created or removed because of insufficient users
// print the name!
// ( The functions need to be of the type func(string) and receive the rooms name as argument )


func main() {
	
        addr , err := determineListenAddress()
         if err != nil {
       log.Fatal(err)
       }    
	// Create a router
       ro := mux.NewRouter()
      
        hub := newHub()
        ou3 := &UniqueRand{} 
        Connesi := ou3.generated
        if Connesi == nil {
          Connesi = make(map[int]bool)
	        ou3.generated = Connesi
        }
        ou3.generated[12] = true 
    //go hub.run()

        hub.AddFirstRoomis()
        
	// Add the events to the router
	ro.HandleFunc("/{team}/channels/{room}/" , hub.LoginQuat())  
        staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	ro.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")


        
        dbUrl := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("[x] Could not open the connection to the database. Reason: %s", err.Error())
	}
        InitStore(&dbStore{db: db})
rmajis, errtoi := store.GetEmojisAll()
log.Println("erreur", errtoi)
log.Println("erreur", rmajis)
count, zert := store.GetCountRooms()
log.Println("rooom", zert)
log.Println("coun", count)

        paii, poiu, ertyui := store.GetTeamsAlla()
log.Println("rooom", ertyui)
log.Println("rooom", paii)
log.Println("rooom", poiu)


        ro.HandleFunc("/{team}/channels/{room}", func(w http.ResponseWriter, r *http.Request) {
                



		if ws, err := NewWebSocket(hub, w, r); err == nil {
                        vars := mux.Vars(r)  
                        log.Println(vars["team"])
                        

                        
                        for _, v := range hub.Teamis {
                        vars := mux.Vars(r) 
                        fmt.Println("Hellotouuttttttttttttt", v)  
              for j, k := range v { 
              fmt.Println("Hello k teams ", k.teams) 
                       for _, kui := range k.teams {
                            fmt.Println("server",kui[0].Name)
                            
                            
                            for _,pio := range kui {
                                fmt.Println("teamtyuuuuuuu", pio)
                                team := pio
                                if vars["team"] == team.Displayname {
                                         fmt.Println("mmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmmm")
                                  } 

                            }
                            
                            


                       }  
              fmt.Println("Hello k rooms", k.rooms) 
              //fmt.Println("Hellonnnnnnnnnn", k.rooms[i.Name]) 
              //fmt.Println("Hello", k.rooms[i.Name][0]) 
              //fmt.Println("Hello", k.rooms[i.Name][1])  
              fmt.Println("Hello 2eme boucle", j)   
                       for o, ku := range k.rooms {
                            vars := mux.Vars(r) 
                            fmt.Println("team2222",o)  /// team
                            for _, kug := range ku {
                                fmt.Println("i m a room",kug)
                                room := kug
                                
                                if vars["room"] == room.Displayname && appTeam.Displayname == vars["team"]{
                                      fmt.Println("i am THEEEEEEEEEEEEEEEEEEEEEEEEE room", vars["room"])
                                      log.Println("appTeam", appTeam)
                                      hub.Register(room, ws)  
                                 

                                
                                 
                         			addr := strings.Split(r.RemoteAddr, ":")[0]

			loginUser := func(uname string) {
                                room.Author = &Author{
					ID:       ou3.Int(),
					Color:    UtilGetRandomColor(),
					Username: uname,
                                        Room: room,  
				}
                                hub.Sackets[ws] = &Author{
					ID:       ou3.Int(),
					Color:    UtilGetRandomColor(),
					Username: uname,
				}


                               


				hub.Sockets[ws] = room

				hub.Loginnn(ws, hub.Sockets[ws].Room.Author, addr)
				hub.Users[hub.Sockets[ws].Room.Author.Username] = hub.Sockets[ws].Room.Author
	                        hub.AppendHistoryPosts(room)
                                log.Println("usersiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii", hub.Teamis["mykeys"]["key"])
                                log.Println("users", hub.Users)
                                
				go func() {
					ws.send <- (&Event{
						Name: "clientConnect",
						Data: map[string]interface{}{
							"author":   hub.Sockets[ws],
							"nclients": len(hub.Sockets),
							"clients":  hub.Users,
                                                        "test": hub.GetTeamis(paii, poiu),
                                                        "tosts": paii,
                                                        "tasto": hub.GetUrls(paii, poiu), 
                                                        "pals": vars["team"], 
                                                        "palis": room, 
                                                        "emo": rmajis, 
                                                        "roomteams": poiu,  
                                                        "hposts":  hub.Posts,   
							"history":  hub.History,
						},
					}).Raw()
				}()
				hub.Broadcast((&Event{
					Name: "connected",
					Data: map[string]interface{}{
						"author":   hub.Sockets[ws],
                                                "room": room,  
                                                "nclients": len(hub.Sockets),
						"clients":  hub.Users,
					},
				}).Raw(), (room ) )
			}

			

			// USERNAME INPUT EVENT
			// -> Checks if name is not connected yet
			//    -> else send 'connect_reject' event
			// -> Broadcast to all clients that user has connected
			ws.SetHandler("login", func(event *Event) {
				dataMap := event.Data.(map[string]interface{})
				uname := dataMap["username"].(string)
				passwd := dataMap["password"].(string)
                                log.Println("pass", passwd)
				
				loginUser(uname)
			})

                        laginUser := func(uname string) {
                                log.Println("pass", uname)
                        }

                        ws.SetHandler("suscribe", func(event *Event) {
				dataMap := event.Data.(map[string]interface{})
				uname := dataMap["username"].(string)
				//log.Println("ccc", req.name) 
                                log.Println("pass", uname)
				log.Println("passoooooooooo", hub.Sackets[ws].Username)
                                pi, _ := store.GetTeamMemberByEmail(hub.Sackets[ws].Username)
                                log.Println("passoooooooooo", pi.Roles)
                                log.Println("passoooooooooo", pi.Userid)
                                s := strconv.FormatInt(count + 1, 10)
                                room := &Room{Id: s , Name: uname, Teamid: pi.Teamid , Displayname: uname , Authorid: pi.Userid } 
                                log.Println("passoooooooooo", pi.Teamid)
                                rat := store.CreateRoomWithUser(room , pi.Userid, pi.Teamid) 
                                log.Println("passoooooooooo", rat)


				laginUser(uname)
			})

			// CHAT MESSAGE EVENT
			// -> Attach username to message
			// -> Broadcast the chat message to all users
			ws.SetHandler("message", func(event *Event) {
				if len(strings.Trim(event.Data.(string), " \t")) < 1 {
					return
				}
				author := hub.Sackets[ws]
                                //m := int64(author.ID)
                                //if hub.TempHistoryLength(m) > 10 {
					//go func() {
						//ws.Out <- (&Event{
							//Name: "spamTimeout",
							//Data: nil,
						//}).Raw()
					//}()
					//return
			        //}
                         

                                mass := &Message{ Content:   strings.Replace(event.Data.(string), "\\n", "\n", -1),  }
    bi := []byte(event.Data.(string))  
    
       for _, b := range bi {
           if escaped(b) > -1 {
              
             //return bi
           } else {
            log.Println("ibbiiiiiii", b)
             //y = append(y, b)
             //return y
           }
     }
    


log.Println("authoriiiiiiiiiiiiiiiiiiiiiiiiiiiiii", event.Data.(string))
log.Println("authoriiiiiiiiiiiiiiiiiiiiiiiiiiiiii", testEsc(event.Data.(string)))
			        
				event.Data = &Message{
					Author:    author.Username,
                                        Color:     author.Color,
                                        Room:      room.Displayname,
					Content:   mass.Content,
					Timestamp: time.Now().Unix(),
					Id:        autoId(),
				}
                                //log.Println("ggggg", room)
				hub.Broadcast(event.Raw(), ( room ))
post := &Post{Id: autoId(), Createat: time.Now().Unix(), Userid: author.Username, Channelid: room.Id, Rootid: strconv.Itoa(ou3.Int()), Message: mass.Content, Color: author.Color, Typo: "message" , Props: &Props{
        Username: author.Username,
        Markdown: mass.Content,
        Channels: struct {
            Channelid string `json:"channelid,omitempty"`
            
        }{
            Channelid:   room.Id,
            
        },
    } , }   
    
        heelo := store.CreatePost(post)
log.Println("post", post)
log.Println("posti", heelo)

                                 
				//hub.AppendHistory(event)
                                //mi := int64(author.ID) 
				//hub.EnqueueTempHistory(mi)
			})

			

			// DISCONNECT EVENT
			// -> Broadcast to all clients that
			//    user has disconnected
			ws.SetHandler("disconnected", func(event *Event) {
				dataMap := event.Data.(map[string]interface{})
				uname := dataMap["name"].(string)
				delete(hub.Users, uname)
				dataMap["clients"] = hub.Users
				event.Data = dataMap
				hub.Broadcast(event.Raw(), (room ))
			})








                              if vars["room"] == room.Displayname {

                              }

                            }
                        }   

                    }          

              } 















                        }
			
		}


	})

        //tmpl := template.Must(template.ParseFiles("public/index.html")) 
        //tmpl.Execute(w , story)
	// Listen
     
	fmt.Println("Listening on", addr)
	http.ListenAndServe(addr, ro)
}
