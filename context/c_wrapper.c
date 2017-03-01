int start_calculation(int * ctx, float * data, int n);
int * get_context(void);
void set(int i);
void sent(int * ctx, int i[], int j, int k);

int start_calculation_c(int * ctx, float * data, int n)
{
    return start_calculation(ctx, data, n);
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
