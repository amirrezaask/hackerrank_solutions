# Jumping on the clouds
#
defmodule CloudWatch do
  def find_path([_cloud], current, path) do
    [current | path]
  end

  def find_path([], _current, path) do
    nil
  end

  def find_path(["1" | _rest], _current, _path) do
    nil
  end

  def find_path(clouds, current, path) do
    shortest_path(
      find_path(Enum.slice(clouds, 1..(length(clouds) - 1)), current + 1, [current | path]),
      find_path(Enum.slice(clouds, 2..(length(clouds) - 1)), current + 2, [current | path])
    )
  end

  def shortest_path(nil, path2) do
    path2
  end

  def shortest_path(path1, nil) do
    path1
  end

  def shortest_path(path1, path2) when length(path1) == 0 do
    path2
  end

  def shortest_path(path1, path2) when length(path2) == 0 do
    path1
  end

  def shortest_path(path1, path2) when length(path1) <= length(path2) do
    path1
  end

  def shortest_path(path1, path2) when length(path2) < length(path1) do
    path2
  end

  def final_answer(path) do
    length(path) - 1
  end
end

IO.read(:stdio, :line)

IO.read(:stdio, :line)
|> String.split()
|> CloudWatch.find_path(0, [])
|> CloudWatch.final_answer()
|> IO.puts()
