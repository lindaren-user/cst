export async function load({ fetch }) {
    let data;
    try {
        const res = await fetch('/api/getUserSession');
        
        if (!res.ok) {
            throw new Error("服务器出错");
        }

        data = await res.json();

        if (data && data.status !== 0) {
            throw new Error("会话不存在");
        }
    } catch (e) {
        data = { account: '', role: '' };
    }
    return data;
}