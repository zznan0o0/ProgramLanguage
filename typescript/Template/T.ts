function tt <A,B>(a: A, b: B): B{
    console.log(a)
    console.log(b)
    return b;
}

tt<string,number>("123", 456)
