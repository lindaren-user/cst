import { redirect } from "@sveltejs/kit";

export async function load({ fetch }) {
    let data;
    try {
        const res = await fetch('/api/getUserSession');
        
        if (!res.ok) {
            throw new Error("服务器出错");
        }

        data = await res.json();

        if (data && data.status !== 0) {
            throw new Error("会话失效");
        }
        
        if (data.role !== 'admin') {
            throw redirect(303, '/'); // 当服务器收到请求并成功处理后，返回一个 303 状态码，告诉客户端去请求一个新的资源
        }
    } catch (e) {
        throw redirect(303, '/error');
    }
}