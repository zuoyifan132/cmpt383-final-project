import inputParser		# SWIG C++ interface
import random
import json
import zmq


# print the chessboard
def graph(l):
	for i in range(3):
		print (l[i], end = " ")
	print ()
	for i in range(3,6):
		print (l[i], end = " ")
	print ()
	for i in range(6,9):
		print (l[i], end = " ") 
	print ()


# could use graph search dfs
def computer_win(l):
	if (l[0]== "x") & (l[1]== "x") & (l[2]== "x"):
		return True
	elif (l[3]== "x") & (l[4]== "x") & (l[5]== "x"):
		return True
	elif (l[6]== "x") & (l[7]== "x") & (l[8]== "x"):
		return True
	elif (l[0]== "x") & (l[3]== "x") & (l[6]== "x"):
		return True
	elif (l[1]== "x") & (l[4]== "x") & (l[7]== "x"):
		return True
	elif (l[2]== "x") & (l[5]== "x") & (l[8]== "x"):
		return True
	elif (l[0]== "x") & (l[4]== "x") & (l[8]== "x"):
		return True
	elif (l[2]== "x") & (l[4]== "x") & (l[6]== "x"):
		return True
	else:
		return False


# could use graph search dfs
def player_win(l):
	if (l[0]=="o") & (l[1]=="o") & (l[2]=="o"):
		return True
	elif (l[3]=="o") & (l[4]=="o") & (l[5]=="o"):
		return True
	elif (l[6]=="o") & (l[7]=="o") & (l[8]=="o"):
		return True
	elif (l[0]=="o") & (l[3]=="o") & (l[6]=="o"):
		return True
	elif (l[1]=="o") & (l[4]=="o") & (l[7]=="o"):
		return True
	elif (l[2]=="o") & (l[5]=="o") & (l[8]=="o"):
		return True
	elif (l[0]=="o") & (l[4]=="o") & (l[8]=="o"):
		return True
	elif (l[2]=="o") & (l[4]=="o") & (l[6]=="o"):
		return True
	else:
		return False


# check if the status of game is draw
# return true if the game is draw, otherwsie retuen false
def draw(l):
	if (player_win(l) == False) & (computer_win(l) == False):
		count = 0
		for i in range(9):
			if (l[i] == "x") or (l[i] == "o"):
				count += 1
		if count == 9:
			return True
		
	else:
		return False


# check if the game is over for player side
# return true if the game is over, otherwise return false
def player_end_game(l):
	if player_win(l) == True:
		print ("Win")
		return True
	elif computer_win(l) == True:
		print ("Lose")
		return True
	elif draw(l):
		print ("Draw")
		return True
	else:
		return False


# check if layer enter us valid and give helper information
# return true is player input is valid, otherwsie return false
def player_enter(l,n):
	if n < '0' or n > '8':
		print ("please enter number only")
		return False
	n = int(n)
	if (n > 8) or (n < 0):
		print ("input not valid please try again")
		return False
	elif (l[n] == "x") or (l[n] == "o"):
		print ("input position is already taken, please try again")
		return False
	else:
		return True


# play a new game 
def AI_first_new_game(number_random_playout, thread):
	# computer "x" player "o"

	# computer first
	print ("*---------------------------game start---------------------------*")
	print()
	print ("AI go first: (computer: x player: o)")
	print("number are available position")
	print ()
	l = "012345678"

	while(True):

		# send request to go end 
		request = json.dumps(l).encode('utf8')
		socket.send(request)

		response_bytes = socket.recv()							# recive the response b
		response = json.loads(response_bytes.decode('utf-8'))	# decode to json int
		print("AI put on position: ", response)

		# pMCTS return invalid result
		if response == -1:
			print("pMCTS error")
			print("restart the game")
			break

		l = l.replace(l[response], "x", 1)

		# show the game board
		graph(l)

		# check after AI enter if game over
		if player_end_game(l) == True:
			break
	
		# player enter
		player_value = 0
		while(True):
			n = input("Enter:")
			if player_enter(l,n) == True:
				player_value = int(n)
				break

		#change chessboard
		l = l.replace(l[player_value], "o", 1)

		# show the game board
		graph(l)

		# check after player enter if game over 
		if player_end_game(l) == True:
			break

		print()
		print ("*----------------------------------------------------------------*")


def Player_first_new_game(number_random_playout, thread):
	# computer "x" player "o"

	# player first
	print ("*---------------------------game start---------------------------*")
	print()
	print ("Player go first: (computer: x player: o)")
	print("number are available position")
	print ()
	l = "012345678"

	graph(l)

	while(True):
		# player enter
		player_value = 0
		while(True):
			n = input("Enter:")
			if player_enter(l,n) == True:
				player_value = int(n)
				break

		#change chessboard
		l = l.replace(l[player_value], "o", 1)

		# show the game board
		#graph(l)
		print()
		print ("*----------------------------------------------------------------*")

		# check after player enter if game over 
		if player_end_game(l) == True:
			break

		# send request to go end 
		request = json.dumps(l).encode('utf8')
		socket.send(request)

		response_bytes = socket.recv()							# recive the response b
		response = json.loads(response_bytes.decode('utf-8'))	# decode to json int
		print("AI put on position: ", response)

		# pMCTS return invalid result
		if response == -1:
			print("pMCTS error")
			print("restart the game")
			break

		l = l.replace(l[response], "x", 1)

		# show the game board
		graph(l)

		# check after AI enter if game over
		if player_end_game(l) == True:
			break


#main
if __name__ == '__main__':

	Address = "tcp://localhost:5555"

	#--------------------------------socket setup--------------------------------#
	print("Requester setup...")

	context = zmq.Context()					# get ZeroMQ context 
	socket = context.socket(zmq.REQ)		# socket zmq.REQ will block on send unless it has successfully received a reply back.
	socket.connect(Address)					# connect to port 5555

	print("Requester setup done\n")

	#--------------------------------handle input--------------------------------#
	
	first = "P"
	mode = "h"
	thread = 4

	# parse the user input command  
	while True:
		command = input("Please input command: (type `help' to get more information) ")
		parsed = inputParser.parser(command)
		# valid input
		if parsed != "" and parsed != "help":
			first = parsed[0]
			mode = parsed[1]
			thread = int(parsed[2:])
			break
		if parsed != "help":
			print("input command not valid, please try again")

	number_random_playout = 0
	# hard mode: large amout of random playout
	# easy more: small amount of random playout
	if mode == "h":
		number_random_playout = 2000
	elif mode == "m":
		number_random_playout = 5
	else:
		number_random_playout = 1

	# send number_random_playout and thread to go end
	request1 = json.dumps(number_random_playout).encode('utf8')
	socket.send(request1)
	response_bytes1 = socket.recv()						
	ack1 = json.loads(response_bytes1.decode('utf-8'))	# receive acknowledgement

	request2 = json.dumps(thread).encode('utf8')
	socket.send(request2)
	response_bytes2 = socket.recv()						# receive acknowledgement
	ack2 = json.loads(response_bytes2.decode('utf-8'))

	if not ack1 or not ack2:
		print("WARNING: go end didn't recive number_random_playout and thread")

	# start the game
	while True:
		if first == "A":
			AI_first_new_game(number_random_playout, thread)
		else:
			Player_first_new_game(number_random_playout, thread)

		# ask player want to play a new game again
		again = input("Want to play again? [y/n]: ")
		if again == "n":
			break

	#--------------------------------close socket--------------------------------#
	# send shutdown signal to replier
	request = json.dumps("!").encode('utf8')
	socket.send(request)

	print()
	print("close requester")
	socket.close()
	context.term()
	#----------------------------------------------------------------------------#

	












