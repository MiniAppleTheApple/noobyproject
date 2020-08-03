using System;
using System.IO;
using OpenTK;
using OpenTK.Graphics;

namespace cstest
{
    class Program : GameWindow
    {
        public Program(int width, int height, string title) : base(width, height, GraphicsMode.Default, title) { }

        static void Main(string[] args)
        {
            Console.WriteLine("hello");
            Console.WriteLine("hello world"); 
            Console.WriteLine("hello world");
        }
    }
}
