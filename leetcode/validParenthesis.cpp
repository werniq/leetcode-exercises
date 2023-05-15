#include <iostream>

struct Node {
    int value;
    Node *next;

    Node(int _value, Node *_next=nullptr):
        value(_value),
        next(_next)
    {}
};

class Stack {
    Node *head;
    int count;

public:
    // definition of the constructor
    Stack()
    {
        head = nullptr;
        count=0;
    }

    Stack(const Stack& other):head(nullptr)
    {
        if (!other.isEmpty())
        {
            head = nullptr;
            count = 0;
            Node *curr = other.head;
            while (curr != nullptr) {
                push(curr->value);
                curr = curr->next;
            }
        }
    }

    // definition of the destructor
    ~Stack()
    {
        while (!isEmpty())
        {
            pop();
        }
    }

    // isEmpty function checks if the stack is empty by checking if the head pointer is pointing to a null pointer.
    // If the head pointer is pointing to a null pointer, it returns true, indicating that the stack is empty
    bool isEmpty() const
    {
        return head == nullptr;
    }


    // pop
    // first checks if the stack is empty by calling the isEmpty function. If the stack is indeed empty, it prints a message indicating that the stack is empty and cannot perform the pop operation. Then it returns, terminating the function.
    // If the stack is not empty, it creates a temporary pointer temp and assigns it the address of the current top node.
    // It updates the top pointer to point to the next node in the stack, effectively removing the top node.
    // It uses the delete operator to deallocate the memory occupied by the top node pointed to by temp.
    // Finally, it decrements the count variable, which keeps track of the number of elements in the stack.
    void pop()
    {
        if (!isEmpty())
        {
            Node *tmp = head;
            head = head -> next;
            delete tmp;
            count--;
        }
        else
        {
            throw std::runtime_error("Stack is empty");
        }
    }

    // print()
    // Prints the elements of the stack from top to bottom.
    // It first checks if the stack is empty by calling the isEmpty function.
    // If the stack is empty, it prints a message indicating that the stack is empty and returns, terminating the function.
    void print()
    {
        if (isEmpty())
        {
            throw std::runtime_error("Stack is empty");
        }

        Node* curr = head;
        while (curr != nullptr)
        {
            std::cout << curr -> value << " ";
            curr = curr -> next;
        }
        std :: cout << std :: endl;
    }

    // size function returns the number of elements in the stack.
    int size()
    {
        return count;
    }

    int top()
    {
        if (!isEmpty())
        {
            return head -> value;
        }
        else
        {
            throw std::runtime_error("Stack is empty");
        }
    }


    // operator+ function overloads the + operator to allow the use of the + operator to push elements into the stack.
    void operator+(int x)
    {
        push(x);
    }

    // push function inserts a new node at the top of the stack.
    void push(int x)
    {
        Node *tmp = new Node(x);
        tmp -> next = head;
        head = tmp;
        count++;
    }

    int peek() const
    {
        if (!isEmpty())
        {
            return head -> value;
        }
        else
        {
            throw std::runtime_error("Stack is empty");
        }
    }

};

bool validParenthesis(std::string s)
{
    Stack stack;
    char x;

    for (int i=0; i<s.length(); i++)
    {
        if (s[i]=='(' || s[i]=='[' || s[i]=='{')
        {
            stack.push(s[i]);
            continue;
        }
        if (stack.isEmpty())
            return false;
        switch (s[i])
        {
            case ')':
                x = stack.top();
                stack.pop();
                if (x=='{' || x=='[')
                    return false;
                break;
            case '}':
                x = stack.top();
                stack.pop();
                if (x=='(' || x=='[')
                    return false;
                break;
            case ']':
                x = stack.top();
                stack.pop();
                if (x =='(' || x == '{')
                    return false;
                break;
        }
    }
    return (stack.isEmpty());
}



int main()
{
    std::string input1 = "{[()]}";
    std::string input2 = "{[()]";
    std::string input3 = "{[()]}(";
    std::string input4 = "(()())((())(((())(()())(()))))(()";
    std::string input5 = "([])([]{}{())}[])";

    std::cout << input1 << " is " << (validParenthesis(input1) ? "valid" : "invalid") << std::endl;
    std::cout << input2 << " is " << (validParenthesis(input2) ? "valid" : "invalid") << std::endl;
    std::cout << input3 << " is " << (validParenthesis(input3) ? "valid" : "invalid") << std::endl;
    std::cout << input4 << " is " << (validParenthesis(input4) ? "valid" : "invalid") << std::endl;
    std::cout << input5 << " is " << (validParenthesis(input5) ? "valid" : "invalid") << std::endl;

    return 0;
}
