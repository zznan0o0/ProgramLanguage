package orz.kenshin;
import orz.kenshin.utils.*;

/**
 * Hello world!
 *
 */
public class App 
{
    public static void main(String[] args){
        Person1 p1 = new Person1();
        p1.age = 20;
        Person2 p2 = new Person2();
        p2.name = "zs";

        HandleData<Person1, Person2> handleData = new HandleData<Person1, Person2>();

        p1 =  handleData.add(p1, p2);
        System.out.println(p1.name);

    }
}
