class Solution {
public:
    string textfield="A";
    string clipboard="";
    int step=0;
    int minSteps(int n) {
        while(textfield.length()<n)
        {
            if((n-textfield.length()) % textfield.length()==0)
            {
                copyAll();
            }
            paste();
        }
        return step;
    }
    void copyAll()
    {
        clipboard=textfield;
        step++;
    }
    void paste()
    {
        textfield=textfield+clipboard;
        step++;
    }
};
