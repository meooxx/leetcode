function isIsomorphic(s: string, t: string): boolean {

    const Amap: Map<string, string> = new Map()
    const Bmap: Map<string, string> = new Map()
    for (let i = 0; i < s.length; i++) {
        // if (s[i] == t[i]) continue
        if (Amap.has(s[i]) || Bmap.has(t[i])) {
            if (Amap.get(s[i]) !== t[i] || Bmap.get(t[i]) != s[i]) { return false }
        }

        Amap.set(s[i], t[i])
        Bmap.set(t[i], s[i])

    }
    return true
};