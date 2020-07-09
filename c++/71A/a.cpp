#include <stdio.h>
#include <iostream>
#include <string>

using namespace std;

int main()
{
    int nwords;
    cin >> nwords;
    for (int i = 0; i < nwords; i++)
    {
        string word;
        cin >> word;
        if(word.length()>10){
            char f = word[0];
            char l = word[word.length()-1];
            cout << f << word.length()-2 << l <<"\n";
            continue; 
        }
        cout << word << "\n";
    }
}