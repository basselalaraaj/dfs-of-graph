def dfs(V, visited, adj):
    visited.append(V)
    for x in adj[V]:
        if x not in visited:
            dfs(x, visited, adj)


def dfsOfGraph(V, adj):
    visited = []
    for i in range(V):
        if i not in visited:
            dfs(i, visited, adj)
    return visited


V = 5
adj = [[1, 2, 3], [0], [0, 4], [0], [2]]
ans = dfsOfGraph(V, adj)
for i in range(len(ans)):
    print(ans[i], end="")
print()

# 0 1 2 4 3
