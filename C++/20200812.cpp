#include <iostream>
#include <string>
#include <cstring>
using namespace std;

int main()
{
  char cc[80];
  string ss = "ABCD";
  const char* p;

  p = ss.c_str();
  strcpy(cc, ss.c_str());
  cout << p << " " << cc << endl;

  p = ss.data();
  cout << *p << endl;

  ss.copy(cc, 3);
  cc[3] = '\0';
  cout << cc << endl;

  ss.copy(cc, string::npos);
  cc[ss.length()] = '\0';
  cout << cc << endl;
  return 0;
}
