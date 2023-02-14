using Leet;

List<int> list = new List<int>();
Random random = new Random();
for (int i = 0; i < 1000; i++)
{
    list.Add(i);
}

Console.WriteLine(list.IndexOf(99));

Console.WriteLine(Search1(list.ToArray(), 99));
int Search1(int[] nums, int target)
{
    int low = 0;
    int high = nums.Length - 1;
    while (low <= high)
    {
        int middle = (low + high) / 2;
        if (nums[middle] < target)
        {
            low = middle + 1;
        }
        else if (nums[middle] > target)
        {
            high = middle - 1;
        }
        else if (nums[middle] == target)
        {
            return middle;
        }
    }
    return -1;
}


Console.ReadKey();


void Series()
{
    SeriesNums seriesNums = new();
    Console.WriteLine(seriesNums.FibonacciNums(7));
    Console.WriteLine(seriesNums.SumFactorial(100));
}

void Search()
{
    Search search = new();
    List<int> list = new List<int>();
    for (int i = 0; i < 100; i++)
    {
        list.Add(i);
    }
    Console.WriteLine(search.Binary(list.ToArray(), 0, list.Count() - 1, 30));
}

void Sort()
{
    Sort sort = new();
    List<int> ls = new List<int>();
    Random rd = new Random();
    for (int i = 0; i < 20; i++)
    {
        ls.Add(rd.Next(0, 100));
    }
    //冒泡
    //sort.Bubble(ls.ToArray());

    //选择
    //sort.Selection(ls.ToArray());

    //快排
    int[] arr = ls.ToArray();
    sort.Quick(arr, 0, arr.Length - 1);
    //Console.WriteLine(String.Join(",", arr));

    //快排2
    sort.Quick2(ls.ToArray(), 0, ls.Count() - 1);
}