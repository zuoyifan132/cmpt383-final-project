#include "inputParser.h"
#include <iostream>
#include <vector>

using namespace std;


void help() {
	cout<<"cmpt383 tic-tac-toe program: "<<endl;
	cout<<"quick start the game(all parameters are default): `tic'"<<endl;
	cout<<"Not mentioned parameters are default value: AI go first, use 4 thread, hard mode"<<endl;
	cout<<""<<endl;
	cout<<"start with `tic' "<<endl;
	cout<<"following option are available:"<<endl;
	cout<<""<<endl;
	cout<<"    [-f:AI] AI go first"<<endl;
	cout<<"    [-f:user] Player go first"<<endl;
	cout<<"    [-t:x] use x number of thread to do calculation"<<endl;
	cout<<"    [-m:easy] easy mode, AI is stupid"<<endl;
	cout<<"    [-m:medium] medium mode, AI is medium smart"<<endl;
	cout<<"    [-m:hard] hard mode, AI is smart"<<endl;
	cout<<""<<endl;
	cout<<"sample input: "<<endl;
	cout<<"    tic -f:AI -t:4 -m:hard"<<endl;
	cout<<""<<endl;
}

void separator(string input, std::vector<string> *partial) {
	for (int i = 0; i < input.length(); i++) {
		if (input[i] == ' ') {
			int j = i;
			string temp = "";

			while (j+1 < input.length() && input[j+1] != ' ') {
				temp += input[j+1];
				j++;
			}

			partial->push_back(temp);
		}
	}
}

string conatinValid_f(string s) {
	// AI first option 
	if (s == "-f:AI")
		return "A";
	// user first option
	else if (s == "-f:user")
		return "P";
	else
		return "";
}

string conatinValid_t(string s) {
	string res = "";

	// too short
	if (s.length() < 4)
		return "";
	if (s[0] != '-' || s[1] != 't' || s[2] != ':')
		return "";
	// ensure is number
	for (int i = 3; i < s.length(); i++) {
		if (int(s[i]) < 48 || int(s[i]) > 57)
			return "";
		else 
			res += s[i];
	}

	return res;
}

string conatinValid_m(string s) {
	// easy mode option 
	if (s == "-m:easy")
		return "e";
	// user first option
	else if (s == "-m:medium")
		return "m";
	else if (s == "-m:hard")
		return "h";
	else
		return "";
}

std::string parser(std::string input) {
	// default value
	string first = "A";
	string mode = "h";
	string thread = "4";

	if (input == "help") {
		help();
		return "help";
	}
	else if (input == "tic")
		return first + mode + thread;
	else{
		// input length is too short
		if (input.length() < 3){
			cout<<"must start with ``tic'"<<endl;
			return "";
		}
		// length  greater than 3
		else {
			// check the first three letter is tic
			if (input[0] != 't' || input[1] != 'i' || input[2] != 'c') {
				cout<<"must start with ``tic'"<<endl;
				return "";
			}
			// first three letters valid 
			else {
				std::vector<string> partial;
				separator(input, &partial);

				for(int i=0;i<partial.size();i++){
					string f = conatinValid_f(partial[i]);
					string m = conatinValid_m(partial[i]);
					string t = conatinValid_t(partial[i]);
					
					if (f == "" && m == "" && t == "")
						return "";

					// the option is -f
					if (f != "")
						first = f;
					// the option is -m
					if (m != "")
						mode = m;
					// the option is -t
					if (t != "")
						thread = t;
				}

       			return first + mode + thread;
			}
		}
	}
	
	return "";
}




