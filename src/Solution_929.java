package leetcode;

import java.util.Set;
import java.util.TreeSet;

public class Solution_929 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Solution_929 ss = new Solution_929();
		// 输入：["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]
		// 输出：2
		// 解释：实际收到邮件的是 "testemail@leetcode.com" 和 "testemail@lee.tcode.com"。
		String a[] = new String[] { "test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com",
				"testemail+david@lee.tcode.com" };
		System.out.println(ss.numUniqueEmails(a));

		String b[] = new String[] { "fg.r.u.uzj+o.pw@kziczvh.com", "r.cyo.g+d.h+b.ja@tgsg.z.com",
				"fg.r.u.uzj+o.f.d@kziczvh.com", "r.cyo.g+ng.r.iq@tgsg.z.com", "fg.r.u.uzj+lp.k@kziczvh.com",
				"r.cyo.g+n.h.e+n.g@tgsg.z.com", "fg.r.u.uzj+k+p.j@kziczvh.com", "fg.r.u.uzj+w.y+b@kziczvh.com",
				"r.cyo.g+x+d.c+f.t@tgsg.z.com", "r.cyo.g+x+t.y.l.i@tgsg.z.com", "r.cyo.g+brxxi@tgsg.z.com",
				"r.cyo.g+z+dr.k.u@tgsg.z.com", "r.cyo.g+d+l.c.n+g@tgsg.z.com", "fg.r.u.uzj+vq.o@kziczvh.com",
				"fg.r.u.uzj+uzq@kziczvh.com", "fg.r.u.uzj+mvz@kziczvh.com", "fg.r.u.uzj+taj@kziczvh.com",
				"fg.r.u.uzj+fek@kziczvh.com" };
		System.out.println(ss.numUniqueEmails(b));
	}

	public int numUniqueEmails(String[] emails) {
		Set<String> set = new TreeSet<String>();

		for (int i = 0; i < emails.length; i++) {
			// System.out.println("email:" + emails[i]);
			String email[] = emails[i].split("@");
			if (email.length < 2) {
				continue;
			}

			char name[] = email[0].toCharArray();
			StringBuilder sb = new StringBuilder();
			for (int j = 0; j < name.length; j++) {
				if (name[j] == '.') {
					continue;
				} else if (name[j] == '+') {
					break;
				} else {
					sb.append(name[j]);
				}
			}

			sb.append('@');
			sb.append(email[1]);
			// System.out.println(sb.toString());
			set.add(sb.toString());
		}

		return set.size();
	}

}
