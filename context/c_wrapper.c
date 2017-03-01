int init(int * i);
int * get_context(void);
void set(int i);

int initC(int * i)
{
    return init(i);
}

void setC(int i)
{
    return set(i);
}

void sentC(int i[], int j, int k)
{
    return sent(i, j, k);
}

int * GetContextC(void)
{
    return get_context();
}
