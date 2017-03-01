int init(void);
void set(int i);

int initC(void)
{
    return init();
}

void setC(int i)
{
    return set(i);
}

void sentC(int i[], int j, int k)
{
    return sent(i, j, k);
}
