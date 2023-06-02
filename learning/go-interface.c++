// this code is similar to ./go-interface.go file but using c++ with class
#include <iostream>
#include <cmath>

class Shape {
public:
    virtual double area() = 0;
};

class Circle : public Shape {
public:
    Circle(double radius) : m_radius(radius) {}
    double area() override {
        return M_PI * m_radius * m_radius;
    }
private:
    double m_radius;
};

class Rectangle : public Shape {
public:
    Rectangle(double width, double height) : m_width(width), m_height(height) {}
    double area() override {
        return m_width * m_height;
    }
private:
    double m_width;
    double m_height;
};

int main() {
    Shape* s;

    Circle c(2.5);
    s = &c;
    std::cout << s->area() << std::endl; // Output: 19.634954084936208

    Rectangle r(2.0, 3.0);
    s = &r;
    std::cout << s->area() << std::endl; // Output: 6

    return 0;
}
