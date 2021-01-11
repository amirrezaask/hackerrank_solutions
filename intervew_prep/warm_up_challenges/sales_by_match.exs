n = IO.read(:stdio, :line)

ar =
  IO.read(:stdio, :line)
  |> String.split()
  |> Enum.reduce(%{}, fn sock, state ->
    Map.update(state, sock, 1, &(&1 + 1))
  end)
  |> Enum.map(fn {key, val} ->
    {key, floor(val / 2)}
  end)
  |> Enum.reduce(0, fn {_, pair_count}, counter ->
    counter + pair_count
  end)
  |> IO.puts()
