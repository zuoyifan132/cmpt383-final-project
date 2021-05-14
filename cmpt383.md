# <center/>CMPT383 Final Project Tic-Tac-Toe

- **Project Goal:**

  Create a user defined tic-tac-toe game that player vs AI. User could input parameter such as difficulty, how go first and number of thread used to compute random playout by a command. The core part is AI need to do tons of random playout to decide next step. Use golang to create multiple thread to do the random playout which is more efficient than sequential calculation. The user defined diffculty is affected by the number of random playout and time of computation is affected by number of thread used.  

- **The three programming languages** :
  - `C++`: Parsing the user input command, extract information and pass to main_game. Since C++ is statically typed language, using C++ to handle string(user input command) is fast and simple.
  - ``Pyton``: The main program is implement using Python, the code strcuture(brain) is implemented using python because python is easy to code and AI logic is easy coding in python. But the disadvantage is python is somehow slower when doing large amount of calculation.
  - ``Go``: The core part(pMCTS function in worker_replier.go) which need to repetedly call random_playout() function thousands of times is implmented in Go, go is easy to handle concurrency and thread channel. Doing large amount of calculation concurrently in go is fast. 

- **Communication methods**:
  - `SWIG`: Communicating between C++ and Python. Python can call C++ function by SWIG interface.
  - `Zeromq`: Commnuicating between Go and Python. Python end is requester send request and Go end is replier response to the request.

- **Run**:
  - 1: Open terminal 1 on my project and **vagrant up** then **vagrant ssh**.
  - 2: Open terminal 2 on my project and **vagrant up** then **vagrant ssh**.
  - 3: On terminal 1 type **cd project/tic-tac-toe**
  - 4: On terminal 2 type **cd project/tic-tac-toe**
  - 5: On terminal 2 type **swig -c++ -python inputParser.i** 
  - 6: On terminal 2 type **g++ -fPIC -I/usr/include/python3.8 -lstdc++ -shared -o _inputParser.so inputParser_wrap.cxx inputParser.cc**
    - Note: Now the SWIG interface is built.
  - 7: On terminal 1 type **go run worker_replier.go**
  - 8: On terminal 2 type **python3 main_game.py**
    - Note: step 5 and step 6 have to be in ordered.
  - 9: Now the program runs, on terminal 2 the program ask for command, you can simply type **tic** and the game will start with default parameters.
    - Note: You can type **help** on terminal and the program will show following help information.
    - Note: the default is AI go first and hard game mode and use 4 thread to do random playout. 
    - Following option are available and you can use any order you like, the missing option will be default. All error input command are handled and will ask you input again:
      - [-f:AI]  means AI go first
      - [-f:user] means player go first
      - [-t:x] means use x number of thread to do calculation
      - [-m:easy] means easy mode, AI is stupid
      - [-m:medium] means medium mode, AI is medium smart
      - [-m:hard] means hard mode, AI is smart
    - Sample command:
      - **tic -f:AI -t:4 -m:hard**     AI go first, use 4 thread and hard game mode.
      - **tic -m:easy -f:user**        user go first, use 4 thread(default) and easy game mode.
  - 10: Now the game start, simply enter the number from 0 to 8 you want to put on the chess board, if you enter invalid input, it will ask you input again. 
  - 11: After one round of game, program will ask you if would like one more game.type **y** to continue and type **n** to exit the program.
    - Note: if you want to restart program again, make sure programs(main_game.py and worker_repiler.go) on 2 terminal are terminated and do above steps() again.
- **Techniques**:
  - ``Monte Carlo tree search``.
  - ``C++ string processing``
  - ``go channel and goroutine``
  - ``DFS``

