int init(int * i);
int * get_context(void);
void set(int i);
void sent(int * ctx, int i[], int j, int k);

int initC(int * i)
{
    return init(i);
}

void setC(int i)
{
    return set(i);
}

void sentC(int * ctx, int i[], int j, int k)
{
    return sent(ctx, i, j, k);
}

int * GetContextC(void)
{
    return get_context();
}
