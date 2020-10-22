using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace Fibonacci
{
    class Program
    {
        static void Main(string[] args)
        {
            int a, b, limite, i, auxiliar; 
            limite = 10000000; 
            a = 1;
            b = 1; 
            DateTime tiempo1 = DateTime.Now; 
            for (i = 2; i < limite; i++)  
            {
                auxiliar = a;
                a = b; 
                b = auxiliar + a;
                
            }
            DateTime tiempo2 = DateTime.Now;
            Double total = new TimeSpan(tiempo2.Ticks - tiempo1.Ticks).TotalMilliseconds / 1000;
            Console.Write(total);
        }
    }
}