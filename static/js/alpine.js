(() => {
var it = !1,
    ot = !1,
    W = [],
    st = -1;

function Wt(e) {
    Mn(e)
}

function Mn(e) {
    W.includes(e) || W.push(e), Nn()
}

function Gt(e) {
    let t = W.indexOf(e);
    t !== -1 && t > st && W.splice(t, 1)
}

function Nn() {
    !ot && !it && (it = !0, queueMicrotask(kn))
}

function kn() {
    it = !1, ot = !0;
    for (let e = 0; e < W.length; e++) W[e](), st = e;
    W.length = 0, st = -1, ot = !1
}
var T, N, F, ct, at = !0;

function Jt(e) {
    at = !1, e(), at = !0
}

function Yt(e) {
    T = e.reactive, F = e.release, N = t => e.effect(t, {
        scheduler: r => {
            at ? Wt(r) : r()
        }
    }), ct = e.raw
}

function lt(e) {
    N = e
}

function Xt(e) {
    let t = () => {};
    return [n => {
        let i = N(n);
        return e._x_effects || (e._x_effects = new Set, e._x_runEffects = () => {
            e._x_effects.forEach(o => o())
        }), e._x_effects.add(i), t = () => {
            i !== void 0 && (e._x_effects.delete(i), F(i))
        }, i
    }, () => {
        t()
    }]
}

function Se(e, t) {
    let r = !0,
        n, i = N(() => {
            let o = e();
            JSON.stringify(o), r ? n = o : queueMicrotask(() => {
                t(o, n), n = o
            }), r = !1
        });
    return () => F(i)
}
var Zt = [],
    Qt = [],
    er = [];

function tr(e) {
    er.push(e)
}

function re(e, t) {
    typeof t == "function" ? (e._x_cleanups || (e._x_cleanups = []), e._x_cleanups.push(t)) : (t = e, Qt.push(t))
}

function Oe(e) {
    Zt.push(e)
}

function Ce(e, t, r) {
    e._x_attributeCleanups || (e._x_attributeCleanups = {}), e._x_attributeCleanups[t] || (e._x_attributeCleanups[t] = []), e._x_attributeCleanups[t].push(r)
}

function ut(e, t) {
    e._x_attributeCleanups && Object.entries(e._x_attributeCleanups).forEach(([r, n]) => {
        (t === void 0 || t.includes(r)) && (n.forEach(i => i()), delete e._x_attributeCleanups[r])
    })
}

function rr(e) {
    for (e._x_effects?.forEach(Gt); e._x_cleanups?.length;) e._x_cleanups.pop()()
}
var ft = new MutationObserver(ht),
    dt = !1;

function de() {
    ft.observe(document, {
        subtree: !0,
        childList: !0,
        attributes: !0,
        attributeOldValue: !0
    }), dt = !0
}

function pt() {
    Dn(), ft.disconnect(), dt = !1
}
var fe = [];

function Dn() {
    let e = ft.takeRecords();
    fe.push(() => e.length > 0 && ht(e));
    let t = fe.length;
    queueMicrotask(() => {
        if (fe.length === t)
            for (; fe.length > 0;) fe.shift()()
    })
}

function m(e) {
    if (!dt) return e();
    pt();
    let t = e();
    return de(), t
}
var mt = !1,
    Ae = [];

function nr() {
    mt = !0
}

function ir() {
    mt = !1, ht(Ae), Ae = []
}

function ht(e) {
    if (mt) {
        Ae = Ae.concat(e);
        return
    }
    let t = [],
        r = new Set,
        n = new Map,
        i = new Map;
    for (let o = 0; o < e.length; o++)
        if (!e[o].target._x_ignoreMutationObserver && (e[o].type === "childList" && (e[o].removedNodes.forEach(s => {
                s.nodeType === 1 && s._x_marker && r.add(s)
            }), e[o].addedNodes.forEach(s => {
                if (s.nodeType === 1) {
                    if (r.has(s)) {
                        r.delete(s);
                        return
                    }
                    s._x_marker || t.push(s)
                }
            })), e[o].type === "attributes")) {
            let s = e[o].target,
                a = e[o].attributeName,
                c = e[o].oldValue,
                l = () => {
                    n.has(s) || n.set(s, []), n.get(s).push({
                        name: a,
                        value: s.getAttribute(a)
                    })
                },
                u = () => {
                    i.has(s) || i.set(s, []), i.get(s).push(a)
                };
            s.hasAttribute(a) && c === null ? l() : s.hasAttribute(a) ? (u(), l()) : u()
        } i.forEach((o, s) => {
        ut(s, o)
    }), n.forEach((o, s) => {
        Zt.forEach(a => a(s, o))
    });
    for (let o of r) t.some(s => s.contains(o)) || Qt.forEach(s => s(o));
    for (let o of t) o.isConnected && er.forEach(s => s(o));
    t = null, r = null, n = null, i = null
}

function Te(e) {
    return P(k(e))
}

function D(e, t, r) {
    return e._x_dataStack = [t, ...k(r || e)], () => {
        e._x_dataStack = e._x_dataStack.filter(n => n !== t)
    }
}

function k(e) {
    return e._x_dataStack ? e._x_dataStack : typeof ShadowRoot == "function" && e instanceof ShadowRoot ? k(e.host) : e.parentNode ? k(e.parentNode) : []
}

function P(e) {
    return new Proxy({
        objects: e
    }, Pn)
}
var Pn = {
    ownKeys({
        objects: e
    }) {
        return Array.from(new Set(e.flatMap(t => Object.keys(t))))
    },
    has({
        objects: e
    }, t) {
        return t == Symbol.unscopables ? !1 : e.some(r => Object.prototype.hasOwnProperty.call(r, t) || Reflect.has(r, t))
    },
    get({
        objects: e
    }, t, r) {
        return t == "toJSON" ? In : Reflect.get(e.find(n => Reflect.has(n, t)) || {}, t, r)
    },
    set({
        objects: e
    }, t, r, n) {
        let i = e.find(s => Object.prototype.hasOwnProperty.call(s, t)) || e[e.length - 1],
            o = Object.getOwnPropertyDescriptor(i, t);
        return o?.set && o?.get ? o.set.call(n, r) || !0 : Reflect.set(i, t, r)
    }
};

function In() {
    return Reflect.ownKeys(this).reduce((t, r) => (t[r] = Reflect.get(this, r), t), {})
}

function Re(e) {
    let t = n => typeof n == "object" && !Array.isArray(n) && n !== null,
        r = (n, i = "") => {
            Object.entries(Object.getOwnPropertyDescriptors(n)).forEach(([o, {
                value: s,
                enumerable: a
            }]) => {
                if (a === !1 || s === void 0 || typeof s == "object" && s !== null && s.__v_skip) return;
                let c = i === "" ? o : `${i}.${o}`;
                typeof s == "object" && s !== null && s._x_interceptor ? n[o] = s.initialize(e, c, o) : t(s) && s !== n && !(s instanceof Element) && r(s, c)
            })
        };
    return r(e)
}

function Me(e, t = () => {}) {
    let r = {
        initialValue: void 0,
        _x_interceptor: !0,
        initialize(n, i, o) {
            return e(this.initialValue, () => $n(n, i), s => _t(n, i, s), i, o)
        }
    };
    return t(r), n => {
        if (typeof n == "object" && n !== null && n._x_interceptor) {
            let i = r.initialize.bind(r);
            r.initialize = (o, s, a) => {
                let c = n.initialize(o, s, a);
                return r.initialValue = c, i(o, s, a)
            }
        } else r.initialValue = n;
        return r
    }
}

function $n(e, t) {
    return t.split(".").reduce((r, n) => r[n], e)
}

function _t(e, t, r) {
    if (typeof t == "string" && (t = t.split(".")), t.length === 1) e[t[0]] = r;
    else {
        if (t.length === 0) throw error;
        return e[t[0]] || (e[t[0]] = {}), _t(e[t[0]], t.slice(1), r)
    }
}
var or = {};

function y(e, t) {
    or[e] = t
}

function G(e, t) {
    let r = Ln(t);
    return Object.entries(or).forEach(([n, i]) => {
        Object.defineProperty(e, `$${n}`, {
            get() {
                return i(t, r)
            },
            enumerable: !1
        })
    }), e
}

function Ln(e) {
    let [t, r] = gt(e), n = {
        interceptor: Me,
        ...t
    };
    return re(e, r), n
}

function Ne(e, t, r, ...n) {
    try {
        return r(...n)
    } catch (i) {
        ne(i, e, t)
    }
}

function ne(e, t, r = void 0) {
    e = Object.assign(e ?? {
        message: "No error message given."
    }, {
        el: t,
        expression: r
    }), console.warn(`Alpine Expression Error: ${e.message}

${r?'Expression: "'+r+`"

`:""}`, t), setTimeout(() => {
        throw e
    }, 0)
}
var ke = !0;

function De(e) {
    let t = ke;
    ke = !1;
    let r = e();
    return ke = t, r
}

function R(e, t, r = {}) {
    let n;
    return x(e, t)(i => n = i, r), n
}

function x(...e) {
    return sr(...e)
}
var sr = jn;

function ar(e) {
    sr = e
}

function jn(e, t) {
    let r = {};
    G(r, e);
    let n = [r, ...k(e)],
        i = typeof t == "function" ? yt(n, t) : Bn(n, t, e);
    return Ne.bind(null, e, t, i)
}

function yt(e, t) {
    return (r = () => {}, {
        scope: n = {},
        params: i = []
    } = {}) => {
        let o = t.apply(P([n, ...e]), i);
        ie(r, o)
    }
}
var xt = {};

function Fn(e, t) {
    if (xt[e]) return xt[e];
    let r = Object.getPrototypeOf(async function() {}).constructor,
        n = /^[\n\s]*if.*\(.*\)/.test(e.trim()) || /^(let|const)\s/.test(e.trim()) ? `(async()=>{ ${e} })()` : e,
        o = (() => {
            try {
                let s = new r(["__self", "scope"], `with (scope) { __self.result = ${n} }; __self.finished = true; return __self.result;`);
                return Object.defineProperty(s, "name", {
                    value: `[Alpine] ${e}`
                }), s
            } catch (s) {
                return ne(s, t, e), Promise.resolve()
            }
        })();
    return xt[e] = o, o
}

function Bn(e, t, r) {
    let n = Fn(t, r);
    return (i = () => {}, {
        scope: o = {},
        params: s = []
    } = {}) => {
        n.result = void 0, n.finished = !1;
        let a = P([o, ...e]);
        if (typeof n == "function") {
            let c = n(n, a).catch(l => ne(l, r, t));
            n.finished ? (ie(i, n.result, a, s, r), n.result = void 0) : c.then(l => {
                ie(i, l, a, s, r)
            }).catch(l => ne(l, r, t)).finally(() => n.result = void 0)
        }
    }
}

function ie(e, t, r, n, i) {
    if (ke && typeof t == "function") {
        let o = t.apply(r, n);
        o instanceof Promise ? o.then(s => ie(e, s, r, n)).catch(s => ne(s, i, t)) : e(o)
    } else typeof t == "object" && t instanceof Promise ? t.then(o => e(o)) : e(t)
}
var Et = "x-";

function C(e = "") {
    return Et + e
}

function cr(e) {
    Et = e
}
var Pe = {};

function d(e, t) {
    return Pe[e] = t, {
        before(r) {
            if (!Pe[r]) {
                console.warn(String.raw`Cannot find directive \`${r}\`. \`${e}\` will use the default order of execution`);
                return
            }
            let n = J.indexOf(r);
            J.splice(n >= 0 ? n : J.indexOf("DEFAULT"), 0, e)
        }
    }
}

function lr(e) {
    return Object.keys(Pe).includes(e)
}

function me(e, t, r) {
    if (t = Array.from(t), e._x_virtualDirectives) {
        let o = Object.entries(e._x_virtualDirectives).map(([a, c]) => ({
                name: a,
                value: c
            })),
            s = vt(o);
        o = o.map(a => s.find(c => c.name === a.name) ? {
            name: `x-bind:${a.name}`,
            value: `"${a.value}"`
        } : a), t = t.concat(o)
    }
    let n = {};
    return t.map(dr((o, s) => n[o] = s)).filter(mr).map(Kn(n, r)).sort(Hn).map(o => zn(e, o))
}

function vt(e) {
    return Array.from(e).map(dr()).filter(t => !mr(t))
}
var bt = !1,
    pe = new Map,
    ur = Symbol();

function fr(e) {
    bt = !0;
    let t = Symbol();
    ur = t, pe.set(t, []);
    let r = () => {
            for (; pe.get(t).length;) pe.get(t).shift()();
            pe.delete(t)
        },
        n = () => {
            bt = !1, r()
        };
    e(r), n()
}

function gt(e) {
    let t = [],
        r = a => t.push(a),
        [n, i] = Xt(e);
    return t.push(i), [{
        Alpine: K,
        effect: n,
        cleanup: r,
        evaluateLater: x.bind(x, e),
        evaluate: R.bind(R, e)
    }, () => t.forEach(a => a())]
}

function zn(e, t) {
    let r = () => {},
        n = Pe[t.type] || r,
        [i, o] = gt(e);
    Ce(e, t.original, o);
    let s = () => {
        e._x_ignore || e._x_ignoreSelf || (n.inline && n.inline(e, t, i), n = n.bind(n, e, t, i), bt ? pe.get(ur).push(n) : n())
    };
    return s.runCleanups = o, s
}
var Ie = (e, t) => ({
        name: r,
        value: n
    }) => (r.startsWith(e) && (r = r.replace(e, t)), {
        name: r,
        value: n
    }),
    $e = e => e;

function dr(e = () => {}) {
    return ({
        name: t,
        value: r
    }) => {
        let {
            name: n,
            value: i
        } = pr.reduce((o, s) => s(o), {
            name: t,
            value: r
        });
        return n !== t && e(n, t), {
            name: n,
            value: i
        }
    }
}
var pr = [];

function oe(e) {
    pr.push(e)
}

function mr({
    name: e
}) {
    return hr().test(e)
}
var hr = () => new RegExp(`^${Et}([^:^.]+)\\b`);

function Kn(e, t) {
    return ({
        name: r,
        value: n
    }) => {
        let i = r.match(hr()),
            o = r.match(/:([a-zA-Z0-9\-_:]+)/),
            s = r.match(/\.[^.\]]+(?=[^\]]*$)/g) || [],
            a = t || e[r] || r;
        return {
            type: i ? i[1] : null,
            value: o ? o[1] : null,
            modifiers: s.map(c => c.replace(".", "")),
            expression: n,
            original: a
        }
    }
}
var wt = "DEFAULT",
    J = ["ignore", "ref", "data", "id", "anchor", "bind", "init", "for", "model", "modelable", "transition", "show", "if", wt, "teleport"];

function Hn(e, t) {
    let r = J.indexOf(e.type) === -1 ? wt : e.type,
        n = J.indexOf(t.type) === -1 ? wt : t.type;
    return J.indexOf(r) - J.indexOf(n)
}

function Y(e, t, r = {}) {
    e.dispatchEvent(new CustomEvent(t, {
        detail: r,
        bubbles: !0,
        composed: !0,
        cancelable: !0
    }))
}

function I(e, t) {
    if (typeof ShadowRoot == "function" && e instanceof ShadowRoot) {
        Array.from(e.children).forEach(i => I(i, t));
        return
    }
    let r = !1;
    if (t(e, () => r = !0), r) return;
    let n = e.firstElementChild;
    for (; n;) I(n, t, !1), n = n.nextElementSibling
}

function E(e, ...t) {
    console.warn(`Alpine Warning: ${e}`, ...t)
}
var _r = !1;

function gr() {
    _r && E("Alpine has already been initialized on this page. Calling Alpine.start() more than once can cause problems."), _r = !0, document.body || E("Unable to initialize. Trying to load Alpine before `<body>` is available. Did you forget to add `defer` in Alpine's `<script>` tag?"), Y(document, "alpine:init"), Y(document, "alpine:initializing"), de(), tr(t => S(t, I)), re(t => $(t)), Oe((t, r) => {
        me(t, r).forEach(n => n())
    });
    let e = t => !X(t.parentElement, !0);
    Array.from(document.querySelectorAll(br().join(","))).filter(e).forEach(t => {
        S(t)
    }), Y(document, "alpine:initialized"), setTimeout(() => {
        qn()
    })
}
var St = [],
    xr = [];

function yr() {
    return St.map(e => e())
}

function br() {
    return St.concat(xr).map(e => e())
}

function Le(e) {
    St.push(e)
}

function je(e) {
    xr.push(e)
}

function X(e, t = !1) {
    return B(e, r => {
        if ((t ? br() : yr()).some(i => r.matches(i))) return !0
    })
}

function B(e, t) {
    if (e) {
        if (t(e)) return e;
        if (e._x_teleportBack && (e = e._x_teleportBack), !!e.parentElement) return B(e.parentElement, t)
    }
}

function wr(e) {
    return yr().some(t => e.matches(t))
}
var Er = [];

function vr(e) {
    Er.push(e)
}
var Vn = 1;

function S(e, t = I, r = () => {}) {
    B(e, n => n._x_ignore) || fr(() => {
        t(e, (n, i) => {
            n._x_marker || (r(n, i), Er.forEach(o => o(n, i)), me(n, n.attributes).forEach(o => o()), n._x_ignore || (n._x_marker = Vn++), n._x_ignore && i())
        })
    })
}

function $(e, t = I) {
    t(e, r => {
        rr(r), ut(r), delete r._x_marker
    })
}

function qn() {
    [
        ["ui", "dialog", ["[x-dialog], [x-popover]"]],
        ["anchor", "anchor", ["[x-anchor]"]],
        ["sort", "sort", ["[x-sort]"]]
    ].forEach(([t, r, n]) => {
        lr(r) || n.some(i => {
            if (document.querySelector(i)) return E(`found "${i}", but missing ${t} plugin`), !0
        })
    })
}
var At = [],
    Ot = !1;

function se(e = () => {}) {
    return queueMicrotask(() => {
        Ot || setTimeout(() => {
            Fe()
        })
    }), new Promise(t => {
        At.push(() => {
            e(), t()
        })
    })
}

function Fe() {
    for (Ot = !1; At.length;) At.shift()()
}

function Sr() {
    Ot = !0
}

function he(e, t) {
    return Array.isArray(t) ? Ar(e, t.join(" ")) : typeof t == "object" && t !== null ? Un(e, t) : typeof t == "function" ? he(e, t()) : Ar(e, t)
}

function Ar(e, t) {
    let r = o => o.split(" ").filter(Boolean),
        n = o => o.split(" ").filter(s => !e.classList.contains(s)).filter(Boolean),
        i = o => (e.classList.add(...o), () => {
            e.classList.remove(...o)
        });
    return t = t === !0 ? t = "" : t || "", i(n(t))
}

function Un(e, t) {
    let r = a => a.split(" ").filter(Boolean),
        n = Object.entries(t).flatMap(([a, c]) => c ? r(a) : !1).filter(Boolean),
        i = Object.entries(t).flatMap(([a, c]) => c ? !1 : r(a)).filter(Boolean),
        o = [],
        s = [];
    return i.forEach(a => {
        e.classList.contains(a) && (e.classList.remove(a), s.push(a))
    }), n.forEach(a => {
        e.classList.contains(a) || (e.classList.add(a), o.push(a))
    }), () => {
        s.forEach(a => e.classList.add(a)), o.forEach(a => e.classList.remove(a))
    }
}

function Z(e, t) {
    return typeof t == "object" && t !== null ? Wn(e, t) : Gn(e, t)
}

function Wn(e, t) {
    let r = {};
    return Object.entries(t).forEach(([n, i]) => {
        r[n] = e.style[n], n.startsWith("--") || (n = Jn(n)), e.style.setProperty(n, i)
    }), setTimeout(() => {
        e.style.length === 0 && e.removeAttribute("style")
    }), () => {
        Z(e, r)
    }
}

function Gn(e, t) {
    let r = e.getAttribute("style", t);
    return e.setAttribute("style", t), () => {
        e.setAttribute("style", r || "")
    }
}

function Jn(e) {
    return e.replace(/([a-z])([A-Z])/g, "$1-$2").toLowerCase()
}

function _e(e, t = () => {}) {
    let r = !1;
    return function() {
        r ? t.apply(this, arguments) : (r = !0, e.apply(this, arguments))
    }
}
d("transition", (e, {
    value: t,
    modifiers: r,
    expression: n
}, {
    evaluate: i
}) => {
    typeof n == "function" && (n = i(n)), n !== !1 && (!n || typeof n == "boolean" ? Xn(e, r, t) : Yn(e, n, t))
});

function Yn(e, t, r) {
    Or(e, he, ""), {
        enter: i => {
            e._x_transition.enter.during = i
        },
        "enter-start": i => {
            e._x_transition.enter.start = i
        },
        "enter-end": i => {
            e._x_transition.enter.end = i
        },
        leave: i => {
            e._x_transition.leave.during = i
        },
        "leave-start": i => {
            e._x_transition.leave.start = i
        },
        "leave-end": i => {
            e._x_transition.leave.end = i
        }
    } [r](t)
}

function Xn(e, t, r) {
    Or(e, Z);
    let n = !t.includes("in") && !t.includes("out") && !r,
        i = n || t.includes("in") || ["enter"].includes(r),
        o = n || t.includes("out") || ["leave"].includes(r);
    t.includes("in") && !n && (t = t.filter((g, b) => b < t.indexOf("out"))), t.includes("out") && !n && (t = t.filter((g, b) => b > t.indexOf("out")));
    let s = !t.includes("opacity") && !t.includes("scale"),
        a = s || t.includes("opacity"),
        c = s || t.includes("scale"),
        l = a ? 0 : 1,
        u = c ? ge(t, "scale", 95) / 100 : 1,
        p = ge(t, "delay", 0) / 1e3,
        h = ge(t, "origin", "center"),
        w = "opacity, transform",
        z = ge(t, "duration", 150) / 1e3,
        ve = ge(t, "duration", 75) / 1e3,
        f = "cubic-bezier(0.4, 0.0, 0.2, 1)";
    i && (e._x_transition.enter.during = {
        transformOrigin: h,
        transitionDelay: `${p}s`,
        transitionProperty: w,
        transitionDuration: `${z}s`,
        transitionTimingFunction: f
    }, e._x_transition.enter.start = {
        opacity: l,
        transform: `scale(${u})`
    }, e._x_transition.enter.end = {
        opacity: 1,
        transform: "scale(1)"
    }), o && (e._x_transition.leave.during = {
        transformOrigin: h,
        transitionDelay: `${p}s`,
        transitionProperty: w,
        transitionDuration: `${ve}s`,
        transitionTimingFunction: f
    }, e._x_transition.leave.start = {
        opacity: 1,
        transform: "scale(1)"
    }, e._x_transition.leave.end = {
        opacity: l,
        transform: `scale(${u})`
    })
}

function Or(e, t, r = {}) {
    e._x_transition || (e._x_transition = {
        enter: {
            during: r,
            start: r,
            end: r
        },
        leave: {
            during: r,
            start: r,
            end: r
        },
        in(n = () => {}, i = () => {}) {
            Be(e, t, {
                during: this.enter.during,
                start: this.enter.start,
                end: this.enter.end
            }, n, i)
        },
        out(n = () => {}, i = () => {}) {
            Be(e, t, {
                during: this.leave.during,
                start: this.leave.start,
                end: this.leave.end
            }, n, i)
        }
    })
}
window.Element.prototype._x_toggleAndCascadeWithTransitions = function(e, t, r, n) {
    let i = document.visibilityState === "visible" ? requestAnimationFrame : setTimeout,
        o = () => i(r);
    if (t) {
        e._x_transition && (e._x_transition.enter || e._x_transition.leave) ? e._x_transition.enter && (Object.entries(e._x_transition.enter.during).length || Object.entries(e._x_transition.enter.start).length || Object.entries(e._x_transition.enter.end).length) ? e._x_transition.in(r) : o() : e._x_transition ? e._x_transition.in(r) : o();
        return
    }
    e._x_hidePromise = e._x_transition ? new Promise((s, a) => {
        e._x_transition.out(() => {}, () => s(n)), e._x_transitioning && e._x_transitioning.beforeCancel(() => a({
            isFromCancelledTransition: !0
        }))
    }) : Promise.resolve(n), queueMicrotask(() => {
        let s = Cr(e);
        s ? (s._x_hideChildren || (s._x_hideChildren = []), s._x_hideChildren.push(e)) : i(() => {
            let a = c => {
                let l = Promise.all([c._x_hidePromise, ...(c._x_hideChildren || []).map(a)]).then(([u]) => u?.());
                return delete c._x_hidePromise, delete c._x_hideChildren, l
            };
            a(e).catch(c => {
                if (!c.isFromCancelledTransition) throw c
            })
        })
    })
};

function Cr(e) {
    let t = e.parentNode;
    if (t) return t._x_hidePromise ? t : Cr(t)
}

function Be(e, t, {
    during: r,
    start: n,
    end: i
} = {}, o = () => {}, s = () => {}) {
    if (e._x_transitioning && e._x_transitioning.cancel(), Object.keys(r).length === 0 && Object.keys(n).length === 0 && Object.keys(i).length === 0) {
        o(), s();
        return
    }
    let a, c, l;
    Zn(e, {
        start() {
            a = t(e, n)
        },
        during() {
            c = t(e, r)
        },
        before: o,
        end() {
            a(), l = t(e, i)
        },
        after: s,
        cleanup() {
            c(), l()
        }
    })
}

function Zn(e, t) {
    let r, n, i, o = _e(() => {
        m(() => {
            r = !0, n || t.before(), i || (t.end(), Fe()), t.after(), e.isConnected && t.cleanup(), delete e._x_transitioning
        })
    });
    e._x_transitioning = {
        beforeCancels: [],
        beforeCancel(s) {
            this.beforeCancels.push(s)
        },
        cancel: _e(function() {
            for (; this.beforeCancels.length;) this.beforeCancels.shift()();
            o()
        }),
        finish: o
    }, m(() => {
        t.start(), t.during()
    }), Sr(), requestAnimationFrame(() => {
        if (r) return;
        let s = Number(getComputedStyle(e).transitionDuration.replace(/,.*/, "").replace("s", "")) * 1e3,
            a = Number(getComputedStyle(e).transitionDelay.replace(/,.*/, "").replace("s", "")) * 1e3;
        s === 0 && (s = Number(getComputedStyle(e).animationDuration.replace("s", "")) * 1e3), m(() => {
            t.before()
        }), n = !0, requestAnimationFrame(() => {
            r || (m(() => {
                t.end()
            }), Fe(), setTimeout(e._x_transitioning.finish, s + a), i = !0)
        })
    })
}

function ge(e, t, r) {
    if (e.indexOf(t) === -1) return r;
    let n = e[e.indexOf(t) + 1];
    if (!n || t === "scale" && isNaN(n)) return r;
    if (t === "duration" || t === "delay") {
        let i = n.match(/([0-9]+)ms/);
        if (i) return i[1]
    }
    return t === "origin" && ["top", "right", "left", "center", "bottom"].includes(e[e.indexOf(t) + 2]) ? [n, e[e.indexOf(t) + 2]].join(" ") : n
}
var L = !1;

function A(e, t = () => {}) {
    return (...r) => L ? t(...r) : e(...r)
}

function Tr(e) {
    return (...t) => L && e(...t)
}
var Rr = [];

function H(e) {
    Rr.push(e)
}

function Mr(e, t) {
    Rr.forEach(r => r(e, t)), L = !0, kr(() => {
        S(t, (r, n) => {
            n(r, () => {})
        })
    }), L = !1
}
var ze = !1;

function Nr(e, t) {
    t._x_dataStack || (t._x_dataStack = e._x_dataStack), L = !0, ze = !0, kr(() => {
        Qn(t)
    }), L = !1, ze = !1
}

function Qn(e) {
    let t = !1;
    S(e, (n, i) => {
        I(n, (o, s) => {
            if (t && wr(o)) return s();
            t = !0, i(o, s)
        })
    })
}

function kr(e) {
    let t = N;
    lt((r, n) => {
        let i = t(r);
        return F(i), () => {}
    }), e(), lt(t)
}

function xe(e, t, r, n = []) {
    switch (e._x_bindings || (e._x_bindings = T({})), e._x_bindings[t] = r, t = n.includes("camel") ? ai(t) : t, t) {
        case "value":
            ei(e, r);
            break;
        case "style":
            ri(e, r);
            break;
        case "class":
            ti(e, r);
            break;
        case "selected":
        case "checked":
            ni(e, t, r);
            break;
        default:
            Pr(e, t, r);
            break
    }
}

function ei(e, t) {
    if (Ct(e)) e.attributes.value === void 0 && (e.value = t), window.fromModel && (typeof t == "boolean" ? e.checked = ye(e.value) === t : e.checked = Dr(e.value, t));
    else if (Ke(e)) Number.isInteger(t) ? e.value = t : !Array.isArray(t) && typeof t != "boolean" && ![null, void 0].includes(t) ? e.value = String(t) : Array.isArray(t) ? e.checked = t.some(r => Dr(r, e.value)) : e.checked = !!t;
    else if (e.tagName === "SELECT") si(e, t);
    else {
        if (e.value === t) return;
        e.value = t === void 0 ? "" : t
    }
}

function ti(e, t) {
    e._x_undoAddedClasses && e._x_undoAddedClasses(), e._x_undoAddedClasses = he(e, t)
}

function ri(e, t) {
    e._x_undoAddedStyles && e._x_undoAddedStyles(), e._x_undoAddedStyles = Z(e, t)
}

function ni(e, t, r) {
    Pr(e, t, r), oi(e, t, r)
}

function Pr(e, t, r) {
    [null, void 0, !1].includes(r) && li(t) ? e.removeAttribute(t) : (Ir(t) && (r = t), ii(e, t, r))
}

function ii(e, t, r) {
    e.getAttribute(t) != r && e.setAttribute(t, r)
}

function oi(e, t, r) {
    e[t] !== r && (e[t] = r)
}

function si(e, t) {
    let r = [].concat(t).map(n => n + "");
    Array.from(e.options).forEach(n => {
        n.selected = r.includes(n.value)
    })
}

function ai(e) {
    return e.toLowerCase().replace(/-(\w)/g, (t, r) => r.toUpperCase())
}

function Dr(e, t) {
    return e == t
}

function ye(e) {
    return [1, "1", "true", "on", "yes", !0].includes(e) ? !0 : [0, "0", "false", "off", "no", !1].includes(e) ? !1 : e ? Boolean(e) : null
}
var ci = new Set(["allowfullscreen", "async", "autofocus", "autoplay", "checked", "controls", "default", "defer", "disabled", "formnovalidate", "inert", "ismap", "itemscope", "loop", "multiple", "muted", "nomodule", "novalidate", "open", "playsinline", "readonly", "required", "reversed", "selected", "shadowrootclonable", "shadowrootdelegatesfocus", "shadowrootserializable"]);

function Ir(e) {
    return ci.has(e)
}

function li(e) {
    return !["aria-pressed", "aria-checked", "aria-expanded", "aria-selected"].includes(e)
}

function $r(e, t, r) {
    return e._x_bindings && e._x_bindings[t] !== void 0 ? e._x_bindings[t] : jr(e, t, r)
}

function Lr(e, t, r, n = !0) {
    if (e._x_bindings && e._x_bindings[t] !== void 0) return e._x_bindings[t];
    if (e._x_inlineBindings && e._x_inlineBindings[t] !== void 0) {
        let i = e._x_inlineBindings[t];
        return i.extract = n, De(() => R(e, i.expression))
    }
    return jr(e, t, r)
}

function jr(e, t, r) {
    let n = e.getAttribute(t);
    return n === null ? typeof r == "function" ? r() : r : n === "" ? !0 : Ir(t) ? !![t, "true"].includes(n) : n
}

function Ke(e) {
    return e.type === "checkbox" || e.localName === "ui-checkbox" || e.localName === "ui-switch"
}

function Ct(e) {
    return e.type === "radio" || e.localName === "ui-radio"
}

function He(e, t) {
    var r;
    return function() {
        var n = this,
            i = arguments,
            o = function() {
                r = null, e.apply(n, i)
            };
        clearTimeout(r), r = setTimeout(o, t)
    }
}

function Ve(e, t) {
    let r;
    return function() {
        let n = this,
            i = arguments;
        r || (e.apply(n, i), r = !0, setTimeout(() => r = !1, t))
    }
}

function qe({
    get: e,
    set: t
}, {
    get: r,
    set: n
}) {
    let i = !0,
        o, s, a = N(() => {
            let c = e(),
                l = r();
            if (i) n(Tt(c)), i = !1;
            else {
                let u = JSON.stringify(c),
                    p = JSON.stringify(l);
                u !== o ? n(Tt(c)) : u !== p && t(Tt(l))
            }
            o = JSON.stringify(e()), s = JSON.stringify(r())
        });
    return () => {
        F(a)
    }
}

function Tt(e) {
    return typeof e == "object" ? JSON.parse(JSON.stringify(e)) : e
}

function Fr(e) {
    (Array.isArray(e) ? e : [e]).forEach(r => r(K))
}
var Q = {},
    Br = !1;

function zr(e, t) {
    if (Br || (Q = T(Q), Br = !0), t === void 0) return Q[e];
    Q[e] = t, Re(Q[e]), typeof t == "object" && t !== null && t.hasOwnProperty("init") && typeof t.init == "function" && Q[e].init()
}

function Kr() {
    return Q
}
var Hr = {};

function Vr(e, t) {
    let r = typeof t != "function" ? () => t : t;
    return e instanceof Element ? Rt(e, r()) : (Hr[e] = r, () => {})
}

function qr(e) {
    return Object.entries(Hr).forEach(([t, r]) => {
        Object.defineProperty(e, t, {
            get() {
                return (...n) => r(...n)
            }
        })
    }), e
}

function Rt(e, t, r) {
    let n = [];
    for (; n.length;) n.pop()();
    let i = Object.entries(t).map(([s, a]) => ({
            name: s,
            value: a
        })),
        o = vt(i);
    return i = i.map(s => o.find(a => a.name === s.name) ? {
        name: `x-bind:${s.name}`,
        value: `"${s.value}"`
    } : s), me(e, i, r).map(s => {
        n.push(s.runCleanups), s()
    }), () => {
        for (; n.length;) n.pop()()
    }
}
var Ur = {};

function Wr(e, t) {
    Ur[e] = t
}

function Gr(e, t) {
    return Object.entries(Ur).forEach(([r, n]) => {
        Object.defineProperty(e, r, {
            get() {
                return (...i) => n.bind(t)(...i)
            },
            enumerable: !1
        })
    }), e
}
var ui = {
        get reactive() {
            return T
        },
        get release() {
            return F
        },
        get effect() {
            return N
        },
        get raw() {
            return ct
        },
        version: "3.14.9",
        flushAndStopDeferringMutations: ir,
        dontAutoEvaluateFunctions: De,
        disableEffectScheduling: Jt,
        startObservingMutations: de,
        stopObservingMutations: pt,
        setReactivityEngine: Yt,
        onAttributeRemoved: Ce,
        onAttributesAdded: Oe,
        closestDataStack: k,
        skipDuringClone: A,
        onlyDuringClone: Tr,
        addRootSelector: Le,
        addInitSelector: je,
        interceptClone: H,
        addScopeToNode: D,
        deferMutations: nr,
        mapAttributes: oe,
        evaluateLater: x,
        interceptInit: vr,
        setEvaluator: ar,
        mergeProxies: P,
        extractProp: Lr,
        findClosest: B,
        onElRemoved: re,
        closestRoot: X,
        destroyTree: $,
        interceptor: Me,
        transition: Be,
        setStyles: Z,
        mutateDom: m,
        directive: d,
        entangle: qe,
        throttle: Ve,
        debounce: He,
        evaluate: R,
        initTree: S,
        nextTick: se,
        prefixed: C,
        prefix: cr,
        plugin: Fr,
        magic: y,
        store: zr,
        start: gr,
        clone: Nr,
        cloneNode: Mr,
        bound: $r,
        $data: Te,
        watch: Se,
        walk: I,
        data: Wr,
        bind: Vr
    },
    K = ui;

function Jr(e, t) {
    let r = fi(e);
    if (typeof t == "function") return yt(r, t);
    let n = di(e, t, r);
    return Ne.bind(null, e, t, n)
}

function fi(e) {
    let t = {};
    return G(t, e), [t, ...k(e)]
}

function di(e, t, r) {
    return (n = () => {}, {
        scope: i = {},
        params: o = []
    } = {}) => {
        let s = P([i, ...r]),
            a = t.split(".").reduce((c, l) => (c[l] === void 0 && pi(e, t), c[l]), s);
        ie(n, a, s, o)
    }
}

function pi(e, t) {
    console.warn(`Alpine Error: Alpine is unable to interpret the following expression using the CSP-friendly build:

"${t}"

Read more about the Alpine's CSP-friendly build restrictions here: https://alpinejs.dev/advanced/csp

`, e)
}

function Mt(e, t) {
    let r = Object.create(null),
        n = e.split(",");
    for (let i = 0; i < n.length; i++) r[n[i]] = !0;
    return t ? i => !!r[i.toLowerCase()] : i => !!r[i]
}
var mi = "itemscope,allowfullscreen,formnovalidate,ismap,nomodule,novalidate,readonly";
var qs = Mt(mi + ",async,autofocus,autoplay,controls,default,defer,disabled,hidden,loop,open,required,reversed,scoped,seamless,checked,muted,multiple,selected");
var Yr = Object.freeze({}),
    Us = Object.freeze([]);
var hi = Object.prototype.hasOwnProperty,
    be = (e, t) => hi.call(e, t),
    V = Array.isArray,
    ae = e => Xr(e) === "[object Map]";
var _i = e => typeof e == "string",
    Ue = e => typeof e == "symbol",
    we = e => e !== null && typeof e == "object";
var gi = Object.prototype.toString,
    Xr = e => gi.call(e),
    Nt = e => Xr(e).slice(8, -1);
var We = e => _i(e) && e !== "NaN" && e[0] !== "-" && "" + parseInt(e, 10) === e;
var Ge = e => {
        let t = Object.create(null);
        return r => t[r] || (t[r] = e(r))
    },
    xi = /-(\w)/g,
    Ws = Ge(e => e.replace(xi, (t, r) => r ? r.toUpperCase() : "")),
    yi = /\B([A-Z])/g,
    Gs = Ge(e => e.replace(yi, "-$1").toLowerCase()),
    kt = Ge(e => e.charAt(0).toUpperCase() + e.slice(1)),
    Js = Ge(e => e ? `on${kt(e)}` : ""),
    Dt = (e, t) => e !== t && (e === e || t === t);
var Pt = new WeakMap,
    Ee = [],
    j, ee = Symbol("iterate"),
    It = Symbol("Map key iterate");

function bi(e) {
    return e && e._isEffect === !0
}

function nn(e, t = Yr) {
    bi(e) && (e = e.raw);
    let r = Ei(e, t);
    return t.lazy || r(), r
}

function on(e) {
    e.active && (sn(e), e.options.onStop && e.options.onStop(), e.active = !1)
}
var wi = 0;

function Ei(e, t) {
    let r = function() {
        if (!r.active) return e();
        if (!Ee.includes(r)) {
            sn(r);
            try {
                return Si(), Ee.push(r), j = r, e()
            } finally {
                Ee.pop(), an(), j = Ee[Ee.length - 1]
            }
        }
    };
    return r.id = wi++, r.allowRecurse = !!t.allowRecurse, r._isEffect = !0, r.active = !0, r.raw = e, r.deps = [], r.options = t, r
}

function sn(e) {
    let {
        deps: t
    } = e;
    if (t.length) {
        for (let r = 0; r < t.length; r++) t[r].delete(e);
        t.length = 0
    }
}
var ce = !0,
    Lt = [];

function vi() {
    Lt.push(ce), ce = !1
}

function Si() {
    Lt.push(ce), ce = !0
}

function an() {
    let e = Lt.pop();
    ce = e === void 0 ? !0 : e
}

function M(e, t, r) {
    if (!ce || j === void 0) return;
    let n = Pt.get(e);
    n || Pt.set(e, n = new Map);
    let i = n.get(r);
    i || n.set(r, i = new Set), i.has(j) || (i.add(j), j.deps.push(i), j.options.onTrack && j.options.onTrack({
        effect: j,
        target: e,
        type: t,
        key: r
    }))
}

function U(e, t, r, n, i, o) {
    let s = Pt.get(e);
    if (!s) return;
    let a = new Set,
        c = u => {
            u && u.forEach(p => {
                (p !== j || p.allowRecurse) && a.add(p)
            })
        };
    if (t === "clear") s.forEach(c);
    else if (r === "length" && V(e)) s.forEach((u, p) => {
        (p === "length" || p >= n) && c(u)
    });
    else switch (r !== void 0 && c(s.get(r)), t) {
        case "add":
            V(e) ? We(r) && c(s.get("length")) : (c(s.get(ee)), ae(e) && c(s.get(It)));
            break;
        case "delete":
            V(e) || (c(s.get(ee)), ae(e) && c(s.get(It)));
            break;
        case "set":
            ae(e) && c(s.get(ee));
            break
    }
    let l = u => {
        u.options.onTrigger && u.options.onTrigger({
            effect: u,
            target: e,
            key: r,
            type: t,
            newValue: n,
            oldValue: i,
            oldTarget: o
        }), u.options.scheduler ? u.options.scheduler(u) : u()
    };
    a.forEach(l)
}
var Ai = Mt("__proto__,__v_isRef,__isVue"),
    cn = new Set(Object.getOwnPropertyNames(Symbol).map(e => Symbol[e]).filter(Ue)),
    Oi = ln();
var Ci = ln(!0);
var Zr = Ti();

function Ti() {
    let e = {};
    return ["includes", "indexOf", "lastIndexOf"].forEach(t => {
        e[t] = function(...r) {
            let n = _(this);
            for (let o = 0, s = this.length; o < s; o++) M(n, "get", o + "");
            let i = n[t](...r);
            return i === -1 || i === !1 ? n[t](...r.map(_)) : i
        }
    }), ["push", "pop", "shift", "unshift", "splice"].forEach(t => {
        e[t] = function(...r) {
            vi();
            let n = _(this)[t].apply(this, r);
            return an(), n
        }
    }), e
}

function ln(e = !1, t = !1) {
    return function(n, i, o) {
        if (i === "__v_isReactive") return !e;
        if (i === "__v_isReadonly") return e;
        if (i === "__v_raw" && o === (e ? t ? Vi : pn : t ? Hi : dn).get(n)) return n;
        let s = V(n);
        if (!e && s && be(Zr, i)) return Reflect.get(Zr, i, o);
        let a = Reflect.get(n, i, o);
        return (Ue(i) ? cn.has(i) : Ai(i)) || (e || M(n, "get", i), t) ? a : $t(a) ? !s || !We(i) ? a.value : a : we(a) ? e ? mn(a) : tt(a) : a
    }
}
var Ri = Mi();

function Mi(e = !1) {
    return function(r, n, i, o) {
        let s = r[n];
        if (!e && (i = _(i), s = _(s), !V(r) && $t(s) && !$t(i))) return s.value = i, !0;
        let a = V(r) && We(n) ? Number(n) < r.length : be(r, n),
            c = Reflect.set(r, n, i, o);
        return r === _(o) && (a ? Dt(i, s) && U(r, "set", n, i, s) : U(r, "add", n, i)), c
    }
}

function Ni(e, t) {
    let r = be(e, t),
        n = e[t],
        i = Reflect.deleteProperty(e, t);
    return i && r && U(e, "delete", t, void 0, n), i
}

function ki(e, t) {
    let r = Reflect.has(e, t);
    return (!Ue(t) || !cn.has(t)) && M(e, "has", t), r
}

function Di(e) {
    return M(e, "iterate", V(e) ? "length" : ee), Reflect.ownKeys(e)
}
var Pi = {
        get: Oi,
        set: Ri,
        deleteProperty: Ni,
        has: ki,
        ownKeys: Di
    },
    Ii = {
        get: Ci,
        set(e, t) {
            return console.warn(`Set operation on key "${String(t)}" failed: target is readonly.`, e), !0
        },
        deleteProperty(e, t) {
            return console.warn(`Delete operation on key "${String(t)}" failed: target is readonly.`, e), !0
        }
    };
var jt = e => we(e) ? tt(e) : e,
    Ft = e => we(e) ? mn(e) : e,
    Bt = e => e,
    et = e => Reflect.getPrototypeOf(e);

function Je(e, t, r = !1, n = !1) {
    e = e.__v_raw;
    let i = _(e),
        o = _(t);
    t !== o && !r && M(i, "get", t), !r && M(i, "get", o);
    let {
        has: s
    } = et(i), a = n ? Bt : r ? Ft : jt;
    if (s.call(i, t)) return a(e.get(t));
    if (s.call(i, o)) return a(e.get(o));
    e !== i && e.get(t)
}

function Ye(e, t = !1) {
    let r = this.__v_raw,
        n = _(r),
        i = _(e);
    return e !== i && !t && M(n, "has", e), !t && M(n, "has", i), e === i ? r.has(e) : r.has(e) || r.has(i)
}

function Xe(e, t = !1) {
    return e = e.__v_raw, !t && M(_(e), "iterate", ee), Reflect.get(e, "size", e)
}

function Qr(e) {
    e = _(e);
    let t = _(this);
    return et(t).has.call(t, e) || (t.add(e), U(t, "add", e, e)), this
}

function en(e, t) {
    t = _(t);
    let r = _(this),
        {
            has: n,
            get: i
        } = et(r),
        o = n.call(r, e);
    o ? fn(r, n, e) : (e = _(e), o = n.call(r, e));
    let s = i.call(r, e);
    return r.set(e, t), o ? Dt(t, s) && U(r, "set", e, t, s) : U(r, "add", e, t), this
}

function tn(e) {
    let t = _(this),
        {
            has: r,
            get: n
        } = et(t),
        i = r.call(t, e);
    i ? fn(t, r, e) : (e = _(e), i = r.call(t, e));
    let o = n ? n.call(t, e) : void 0,
        s = t.delete(e);
    return i && U(t, "delete", e, void 0, o), s
}

function rn() {
    let e = _(this),
        t = e.size !== 0,
        r = ae(e) ? new Map(e) : new Set(e),
        n = e.clear();
    return t && U(e, "clear", void 0, void 0, r), n
}

function Ze(e, t) {
    return function(n, i) {
        let o = this,
            s = o.__v_raw,
            a = _(s),
            c = t ? Bt : e ? Ft : jt;
        return !e && M(a, "iterate", ee), s.forEach((l, u) => n.call(i, c(l), c(u), o))
    }
}

function Qe(e, t, r) {
    return function(...n) {
        let i = this.__v_raw,
            o = _(i),
            s = ae(o),
            a = e === "entries" || e === Symbol.iterator && s,
            c = e === "keys" && s,
            l = i[e](...n),
            u = r ? Bt : t ? Ft : jt;
        return !t && M(o, "iterate", c ? It : ee), {
            next() {
                let {
                    value: p,
                    done: h
                } = l.next();
                return h ? {
                    value: p,
                    done: h
                } : {
                    value: a ? [u(p[0]), u(p[1])] : u(p),
                    done: h
                }
            },
            [Symbol.iterator]() {
                return this
            }
        }
    }
}

function q(e) {
    return function(...t) {
        {
            let r = t[0] ? `on key "${t[0]}" ` : "";
            console.warn(`${kt(e)} operation ${r}failed: target is readonly.`, _(this))
        }
        return e === "delete" ? !1 : this
    }
}

function $i() {
    let e = {
            get(o) {
                return Je(this, o)
            },
            get size() {
                return Xe(this)
            },
            has: Ye,
            add: Qr,
            set: en,
            delete: tn,
            clear: rn,
            forEach: Ze(!1, !1)
        },
        t = {
            get(o) {
                return Je(this, o, !1, !0)
            },
            get size() {
                return Xe(this)
            },
            has: Ye,
            add: Qr,
            set: en,
            delete: tn,
            clear: rn,
            forEach: Ze(!1, !0)
        },
        r = {
            get(o) {
                return Je(this, o, !0)
            },
            get size() {
                return Xe(this, !0)
            },
            has(o) {
                return Ye.call(this, o, !0)
            },
            add: q("add"),
            set: q("set"),
            delete: q("delete"),
            clear: q("clear"),
            forEach: Ze(!0, !1)
        },
        n = {
            get(o) {
                return Je(this, o, !0, !0)
            },
            get size() {
                return Xe(this, !0)
            },
            has(o) {
                return Ye.call(this, o, !0)
            },
            add: q("add"),
            set: q("set"),
            delete: q("delete"),
            clear: q("clear"),
            forEach: Ze(!0, !0)
        };
    return ["keys", "values", "entries", Symbol.iterator].forEach(o => {
        e[o] = Qe(o, !1, !1), r[o] = Qe(o, !0, !1), t[o] = Qe(o, !1, !0), n[o] = Qe(o, !0, !0)
    }), [e, r, t, n]
}
var [Li, ji, Fi, Bi] = $i();

function un(e, t) {
    let r = t ? e ? Bi : Fi : e ? ji : Li;
    return (n, i, o) => i === "__v_isReactive" ? !e : i === "__v_isReadonly" ? e : i === "__v_raw" ? n : Reflect.get(be(r, i) && i in n ? r : n, i, o)
}
var zi = {
    get: un(!1, !1)
};
var Ki = {
    get: un(!0, !1)
};

function fn(e, t, r) {
    let n = _(r);
    if (n !== r && t.call(e, n)) {
        let i = Nt(e);
        console.warn(`Reactive ${i} contains both the raw and reactive versions of the same object${i==="Map"?" as keys":""}, which can lead to inconsistencies. Avoid differentiating between the raw and reactive versions of an object and only use the reactive version if possible.`)
    }
}
var dn = new WeakMap,
    Hi = new WeakMap,
    pn = new WeakMap,
    Vi = new WeakMap;

function qi(e) {
    switch (e) {
        case "Object":
        case "Array":
            return 1;
        case "Map":
        case "Set":
        case "WeakMap":
        case "WeakSet":
            return 2;
        default:
            return 0
    }
}

function Ui(e) {
    return e.__v_skip || !Object.isExtensible(e) ? 0 : qi(Nt(e))
}

function tt(e) {
    return e && e.__v_isReadonly ? e : hn(e, !1, Pi, zi, dn)
}

function mn(e) {
    return hn(e, !0, Ii, Ki, pn)
}

function hn(e, t, r, n, i) {
    if (!we(e)) return console.warn(`value cannot be made reactive: ${String(e)}`), e;
    if (e.__v_raw && !(t && e.__v_isReactive)) return e;
    let o = i.get(e);
    if (o) return o;
    let s = Ui(e);
    if (s === 0) return e;
    let a = new Proxy(e, s === 2 ? n : r);
    return i.set(e, a), a
}

function _(e) {
    return e && _(e.__v_raw) || e
}

function $t(e) {
    return Boolean(e && e.__v_isRef === !0)
}
y("nextTick", () => se);
y("dispatch", e => Y.bind(Y, e));
y("watch", (e, {
    evaluateLater: t,
    cleanup: r
}) => (n, i) => {
    let o = t(n),
        a = Se(() => {
            let c;
            return o(l => c = l), c
        }, i);
    r(a)
});
y("store", Kr);
y("data", e => Te(e));
y("root", e => X(e));
y("refs", e => (e._x_refs_proxy || (e._x_refs_proxy = P(Wi(e))), e._x_refs_proxy));

function Wi(e) {
    let t = [];
    return B(e, r => {
        r._x_refs && t.push(r._x_refs)
    }), t
}
var zt = {};

function Kt(e) {
    return zt[e] || (zt[e] = 0), ++zt[e]
}

function _n(e, t) {
    return B(e, r => {
        if (r._x_ids && r._x_ids[t]) return !0
    })
}

function gn(e, t) {
    e._x_ids || (e._x_ids = {}), e._x_ids[t] || (e._x_ids[t] = Kt(t))
}
y("id", (e, {
    cleanup: t
}) => (r, n = null) => {
    let i = `${r}${n?`-${n}`:""}`;
    return Gi(e, i, t, () => {
        let o = _n(e, r),
            s = o ? o._x_ids[r] : Kt(r);
        return n ? `${r}-${s}-${n}` : `${r}-${s}`
    })
});
H((e, t) => {
    e._x_id && (t._x_id = e._x_id)
});

function Gi(e, t, r, n) {
    if (e._x_id || (e._x_id = {}), e._x_id[t]) return e._x_id[t];
    let i = n();
    return e._x_id[t] = i, r(() => {
        delete e._x_id[t]
    }), i
}
y("el", e => e);
xn("Focus", "focus", "focus");
xn("Persist", "persist", "persist");

function xn(e, t, r) {
    y(t, n => E(`You can't use [$${t}] without first installing the "${e}" plugin here: https://alpinejs.dev/plugins/${r}`, n))
}
d("modelable", (e, {
    expression: t
}, {
    effect: r,
    evaluateLater: n,
    cleanup: i
}) => {
    let o = n(t),
        s = () => {
            let u;
            return o(p => u = p), u
        },
        a = n(`${t} = __placeholder`),
        c = u => a(() => {}, {
            scope: {
                __placeholder: u
            }
        }),
        l = s();
    c(l), queueMicrotask(() => {
        if (!e._x_model) return;
        e._x_removeModelListeners.default();
        let u = e._x_model.get,
            p = e._x_model.set,
            h = qe({
                get() {
                    return u()
                },
                set(w) {
                    p(w)
                }
            }, {
                get() {
                    return s()
                },
                set(w) {
                    c(w)
                }
            });
        i(h)
    })
});
d("teleport", (e, {
    modifiers: t,
    expression: r
}, {
    cleanup: n
}) => {
    e.tagName.toLowerCase() !== "template" && E("x-teleport can only be used on a <template> tag", e);
    let i = yn(r),
        o = e.content.cloneNode(!0).firstElementChild;
    e._x_teleport = o, o._x_teleportBack = e, e.setAttribute("data-teleport-template", !0), o.setAttribute("data-teleport-target", !0), e._x_forwardEvents && e._x_forwardEvents.forEach(a => {
        o.addEventListener(a, c => {
            c.stopPropagation(), e.dispatchEvent(new c.constructor(c.type, c))
        })
    }), D(o, {}, e);
    let s = (a, c, l) => {
        l.includes("prepend") ? c.parentNode.insertBefore(a, c) : l.includes("append") ? c.parentNode.insertBefore(a, c.nextSibling) : c.appendChild(a)
    };
    m(() => {
        s(o, i, t), A(() => {
            S(o)
        })()
    }), e._x_teleportPutBack = () => {
        let a = yn(r);
        m(() => {
            s(e._x_teleport, a, t)
        })
    }, n(() => m(() => {
        o.remove(), $(o)
    }))
});
var Ji = document.createElement("div");

function yn(e) {
    let t = A(() => document.querySelector(e), () => Ji)();
    return t || E(`Cannot find x-teleport element for selector: "${e}"`), t
}
var bn = () => {};
bn.inline = (e, {
    modifiers: t
}, {
    cleanup: r
}) => {
    t.includes("self") ? e._x_ignoreSelf = !0 : e._x_ignore = !0, r(() => {
        t.includes("self") ? delete e._x_ignoreSelf : delete e._x_ignore
    })
};
d("ignore", bn);
d("effect", A((e, {
    expression: t
}, {
    effect: r
}) => {
    r(x(e, t))
}));

function le(e, t, r, n) {
    let i = e,
        o = c => n(c),
        s = {},
        a = (c, l) => u => l(c, u);
    if (r.includes("dot") && (t = Yi(t)), r.includes("camel") && (t = Xi(t)), r.includes("passive") && (s.passive = !0), r.includes("capture") && (s.capture = !0), r.includes("window") && (i = window), r.includes("document") && (i = document), r.includes("debounce")) {
        let c = r[r.indexOf("debounce") + 1] || "invalid-wait",
            l = rt(c.split("ms")[0]) ? Number(c.split("ms")[0]) : 250;
        o = He(o, l)
    }
    if (r.includes("throttle")) {
        let c = r[r.indexOf("throttle") + 1] || "invalid-wait",
            l = rt(c.split("ms")[0]) ? Number(c.split("ms")[0]) : 250;
        o = Ve(o, l)
    }
    return r.includes("prevent") && (o = a(o, (c, l) => {
        l.preventDefault(), c(l)
    })), r.includes("stop") && (o = a(o, (c, l) => {
        l.stopPropagation(), c(l)
    })), r.includes("once") && (o = a(o, (c, l) => {
        c(l), i.removeEventListener(t, o, s)
    })), (r.includes("away") || r.includes("outside")) && (i = document, o = a(o, (c, l) => {
        e.contains(l.target) || l.target.isConnected !== !1 && (e.offsetWidth < 1 && e.offsetHeight < 1 || e._x_isShown !== !1 && c(l))
    })), r.includes("self") && (o = a(o, (c, l) => {
        l.target === e && c(l)
    })), (Qi(t) || En(t)) && (o = a(o, (c, l) => {
        eo(l, r) || c(l)
    })), i.addEventListener(t, o, s), () => {
        i.removeEventListener(t, o, s)
    }
}

function Yi(e) {
    return e.replace(/-/g, ".")
}

function Xi(e) {
    return e.toLowerCase().replace(/-(\w)/g, (t, r) => r.toUpperCase())
}

function rt(e) {
    return !Array.isArray(e) && !isNaN(e)
}

function Zi(e) {
    return [" ", "_"].includes(e) ? e : e.replace(/([a-z])([A-Z])/g, "$1-$2").replace(/[_\s]/, "-").toLowerCase()
}

function Qi(e) {
    return ["keydown", "keyup"].includes(e)
}

function En(e) {
    return ["contextmenu", "click", "mouse"].some(t => e.includes(t))
}

function eo(e, t) {
    let r = t.filter(o => !["window", "document", "prevent", "stop", "once", "capture", "self", "away", "outside", "passive"].includes(o));
    if (r.includes("debounce")) {
        let o = r.indexOf("debounce");
        r.splice(o, rt((r[o + 1] || "invalid-wait").split("ms")[0]) ? 2 : 1)
    }
    if (r.includes("throttle")) {
        let o = r.indexOf("throttle");
        r.splice(o, rt((r[o + 1] || "invalid-wait").split("ms")[0]) ? 2 : 1)
    }
    if (r.length === 0 || r.length === 1 && wn(e.key).includes(r[0])) return !1;
    let i = ["ctrl", "shift", "alt", "meta", "cmd", "super"].filter(o => r.includes(o));
    return r = r.filter(o => !i.includes(o)), !(i.length > 0 && i.filter(s => ((s === "cmd" || s === "super") && (s = "meta"), e[`${s}Key`])).length === i.length && (En(e.type) || wn(e.key).includes(r[0])))
}

function wn(e) {
    if (!e) return [];
    e = Zi(e);
    let t = {
        ctrl: "control",
        slash: "/",
        space: " ",
        spacebar: " ",
        cmd: "meta",
        esc: "escape",
        up: "arrow-up",
        down: "arrow-down",
        left: "arrow-left",
        right: "arrow-right",
        period: ".",
        comma: ",",
        equal: "=",
        minus: "-",
        underscore: "_"
    };
    return t[e] = e, Object.keys(t).map(r => {
        if (t[r] === e) return r
    }).filter(r => r)
}
d("model", (e, {
    modifiers: t,
    expression: r
}, {
    effect: n,
    cleanup: i
}) => {
    let o = e;
    t.includes("parent") && (o = e.parentNode);
    let s = x(o, r),
        a;
    typeof r == "string" ? a = x(o, `${r} = __placeholder`) : typeof r == "function" && typeof r() == "string" ? a = x(o, `${r()} = __placeholder`) : a = () => {};
    let c = () => {
            let h;
            return s(w => h = w), vn(h) ? h.get() : h
        },
        l = h => {
            let w;
            s(z => w = z), vn(w) ? w.set(h) : a(() => {}, {
                scope: {
                    __placeholder: h
                }
            })
        };
    typeof r == "string" && e.type === "radio" && m(() => {
        e.hasAttribute("name") || e.setAttribute("name", r)
    });
    var u = e.tagName.toLowerCase() === "select" || ["checkbox", "radio"].includes(e.type) || t.includes("lazy") ? "change" : "input";
    let p = L ? () => {} : le(e, u, t, h => {
        l(Ht(e, t, h, c()))
    });
    if (t.includes("fill") && ([void 0, null, ""].includes(c()) || Ke(e) && Array.isArray(c()) || e.tagName.toLowerCase() === "select" && e.multiple) && l(Ht(e, t, {
            target: e
        }, c())), e._x_removeModelListeners || (e._x_removeModelListeners = {}), e._x_removeModelListeners.default = p, i(() => e._x_removeModelListeners.default()), e.form) {
        let h = le(e.form, "reset", [], w => {
            se(() => e._x_model && e._x_model.set(Ht(e, t, {
                target: e
            }, c())))
        });
        i(() => h())
    }
    e._x_model = {
        get() {
            return c()
        },
        set(h) {
            l(h)
        }
    }, e._x_forceModelUpdate = h => {
        h === void 0 && typeof r == "string" && r.match(/\./) && (h = ""), window.fromModel = !0, m(() => xe(e, "value", h)), delete window.fromModel
    }, n(() => {
        let h = c();
        t.includes("unintrusive") && document.activeElement.isSameNode(e) || e._x_forceModelUpdate(h)
    })
});

function Ht(e, t, r, n) {
    return m(() => {
        if (r instanceof CustomEvent && r.detail !== void 0) return r.detail !== null && r.detail !== void 0 ? r.detail : r.target.value;
        if (Ke(e))
            if (Array.isArray(n)) {
                let i = null;
                return t.includes("number") ? i = Vt(r.target.value) : t.includes("boolean") ? i = ye(r.target.value) : i = r.target.value, r.target.checked ? n.includes(i) ? n : n.concat([i]) : n.filter(o => !to(o, i))
            } else return r.target.checked;
        else {
            if (e.tagName.toLowerCase() === "select" && e.multiple) return t.includes("number") ? Array.from(r.target.selectedOptions).map(i => {
                let o = i.value || i.text;
                return Vt(o)
            }) : t.includes("boolean") ? Array.from(r.target.selectedOptions).map(i => {
                let o = i.value || i.text;
                return ye(o)
            }) : Array.from(r.target.selectedOptions).map(i => i.value || i.text);
            {
                let i;
                return Ct(e) ? r.target.checked ? i = r.target.value : i = n : i = r.target.value, t.includes("number") ? Vt(i) : t.includes("boolean") ? ye(i) : t.includes("trim") ? i.trim() : i
            }
        }
    })
}

function Vt(e) {
    let t = e ? parseFloat(e) : null;
    return ro(t) ? t : e
}

function to(e, t) {
    return e == t
}

function ro(e) {
    return !Array.isArray(e) && !isNaN(e)
}

function vn(e) {
    return e !== null && typeof e == "object" && typeof e.get == "function" && typeof e.set == "function"
}
d("cloak", e => queueMicrotask(() => m(() => e.removeAttribute(C("cloak")))));
je(() => `[${C("init")}]`);
d("init", A((e, {
    expression: t
}, {
    evaluate: r
}) => typeof t == "string" ? !!t.trim() && r(t, {}, !1) : r(t, {}, !1)));
d("text", (e, {
    expression: t
}, {
    effect: r,
    evaluateLater: n
}) => {
    let i = n(t);
    r(() => {
        i(o => {
            m(() => {
                e.textContent = o
            })
        })
    })
});
d("html", (e, {
    expression: t
}, {
    effect: r,
    evaluateLater: n
}) => {
    let i = n(t);
    r(() => {
        i(o => {
            m(() => {
                e.innerHTML = o, e._x_ignoreSelf = !0, S(e), delete e._x_ignoreSelf
            })
        })
    })
});
oe(Ie(":", $e(C("bind:"))));
var Sn = (e, {
    value: t,
    modifiers: r,
    expression: n,
    original: i
}, {
    effect: o,
    cleanup: s
}) => {
    if (!t) {
        let c = {};
        qr(c), x(e, n)(u => {
            Rt(e, u, i)
        }, {
            scope: c
        });
        return
    }
    if (t === "key") return no(e, n);
    if (e._x_inlineBindings && e._x_inlineBindings[t] && e._x_inlineBindings[t].extract) return;
    let a = x(e, n);
    o(() => a(c => {
        c === void 0 && typeof n == "string" && n.match(/\./) && (c = ""), m(() => xe(e, t, c, r))
    })), s(() => {
        e._x_undoAddedClasses && e._x_undoAddedClasses(), e._x_undoAddedStyles && e._x_undoAddedStyles()
    })
};
Sn.inline = (e, {
    value: t,
    modifiers: r,
    expression: n
}) => {
    t && (e._x_inlineBindings || (e._x_inlineBindings = {}), e._x_inlineBindings[t] = {
        expression: n,
        extract: !1
    })
};
d("bind", Sn);

function no(e, t) {
    e._x_keyExpression = t
}
Le(() => `[${C("data")}]`);
d("data", (e, {
    expression: t
}, {
    cleanup: r
}) => {
    if (io(e)) return;
    t = t === "" ? "{}" : t;
    let n = {};
    G(n, e);
    let i = {};
    Gr(i, n);
    let o = R(e, t, {
        scope: i
    });
    (o === void 0 || o === !0) && (o = {}), G(o, e);
    let s = T(o);
    Re(s);
    let a = D(e, s);
    s.init && R(e, s.init), r(() => {
        s.destroy && R(e, s.destroy), a()
    })
});
H((e, t) => {
    e._x_dataStack && (t._x_dataStack = e._x_dataStack, t.setAttribute("data-has-alpine-state", !0))
});

function io(e) {
    return L ? ze ? !0 : e.hasAttribute("data-has-alpine-state") : !1
}
d("show", (e, {
    modifiers: t,
    expression: r
}, {
    effect: n
}) => {
    let i = x(e, r);
    e._x_doHide || (e._x_doHide = () => {
        m(() => {
            e.style.setProperty("display", "none", t.includes("important") ? "important" : void 0)
        })
    }), e._x_doShow || (e._x_doShow = () => {
        m(() => {
            e.style.length === 1 && e.style.display === "none" ? e.removeAttribute("style") : e.style.removeProperty("display")
        })
    });
    let o = () => {
            e._x_doHide(), e._x_isShown = !1
        },
        s = () => {
            e._x_doShow(), e._x_isShown = !0
        },
        a = () => setTimeout(s),
        c = _e(p => p ? s() : o(), p => {
            typeof e._x_toggleAndCascadeWithTransitions == "function" ? e._x_toggleAndCascadeWithTransitions(e, p, s, o) : p ? a() : o()
        }),
        l, u = !0;
    n(() => i(p => {
        !u && p === l || (t.includes("immediate") && (p ? a() : o()), c(p), l = p, u = !1)
    }))
});
d("for", (e, {
    expression: t
}, {
    effect: r,
    cleanup: n
}) => {
    let i = so(t),
        o = x(e, i.items),
        s = x(e, e._x_keyExpression || "index");
    e._x_prevKeys = [], e._x_lookup = {}, r(() => oo(e, i, o, s)), n(() => {
        Object.values(e._x_lookup).forEach(a => m(() => {
            $(a), a.remove()
        })), delete e._x_prevKeys, delete e._x_lookup
    })
});

function oo(e, t, r, n) {
    let i = s => typeof s == "object" && !Array.isArray(s),
        o = e;
    r(s => {
        ao(s) && s >= 0 && (s = Array.from(Array(s).keys(), f => f + 1)), s === void 0 && (s = []);
        let a = e._x_lookup,
            c = e._x_prevKeys,
            l = [],
            u = [];
        if (i(s)) s = Object.entries(s).map(([f, g]) => {
            let b = An(t, g, f, s);
            n(v => {
                u.includes(v) && E("Duplicate key on x-for", e), u.push(v)
            }, {
                scope: {
                    index: f,
                    ...b
                }
            }), l.push(b)
        });
        else
            for (let f = 0; f < s.length; f++) {
                let g = An(t, s[f], f, s);
                n(b => {
                    u.includes(b) && E("Duplicate key on x-for", e), u.push(b)
                }, {
                    scope: {
                        index: f,
                        ...g
                    }
                }), l.push(g)
            }
        let p = [],
            h = [],
            w = [],
            z = [];
        for (let f = 0; f < c.length; f++) {
            let g = c[f];
            u.indexOf(g) === -1 && w.push(g)
        }
        c = c.filter(f => !w.includes(f));
        let ve = "template";
        for (let f = 0; f < u.length; f++) {
            let g = u[f],
                b = c.indexOf(g);
            if (b === -1) c.splice(f, 0, g), p.push([ve, f]);
            else if (b !== f) {
                let v = c.splice(f, 1)[0],
                    O = c.splice(b - 1, 1)[0];
                c.splice(f, 0, O), c.splice(b, 0, v), h.push([v, O])
            } else z.push(g);
            ve = g
        }
        for (let f = 0; f < w.length; f++) {
            let g = w[f];
            g in a && (m(() => {
                $(a[g]), a[g].remove()
            }), delete a[g])
        }
        for (let f = 0; f < h.length; f++) {
            let [g, b] = h[f], v = a[g], O = a[b], te = document.createElement("div");
            m(() => {
                O || E('x-for ":key" is undefined or invalid', o, b, a), O.after(te), v.after(O), O._x_currentIfEl && O.after(O._x_currentIfEl), te.before(v), v._x_currentIfEl && v.after(v._x_currentIfEl), te.remove()
            }), O._x_refreshXForScope(l[u.indexOf(b)])
        }
        for (let f = 0; f < p.length; f++) {
            let [g, b] = p[f], v = g === "template" ? o : a[g];
            v._x_currentIfEl && (v = v._x_currentIfEl);
            let O = l[b],
                te = u[b],
                ue = document.importNode(o.content, !0).firstElementChild,
                Ut = T(O);
            D(ue, Ut, o), ue._x_refreshXForScope = Cn => {
                Object.entries(Cn).forEach(([Tn, Rn]) => {
                    Ut[Tn] = Rn
                })
            }, m(() => {
                v.after(ue), A(() => S(ue))()
            }), typeof te == "object" && E("x-for key cannot be an object, it must be a string or an integer", o), a[te] = ue
        }
        for (let f = 0; f < z.length; f++) a[z[f]]._x_refreshXForScope(l[u.indexOf(z[f])]);
        o._x_prevKeys = u
    })
}

function so(e) {
    let t = /,([^,\}\]]*)(?:,([^,\}\]]*))?$/,
        r = /^\s*\(|\)\s*$/g,
        n = /([\s\S]*?)\s+(?:in|of)\s+([\s\S]*)/,
        i = e.match(n);
    if (!i) return;
    let o = {};
    o.items = i[2].trim();
    let s = i[1].replace(r, "").trim(),
        a = s.match(t);
    return a ? (o.item = s.replace(t, "").trim(), o.index = a[1].trim(), a[2] && (o.collection = a[2].trim())) : o.item = s, o
}

function An(e, t, r, n) {
    let i = {};
    return /^\[.*\]$/.test(e.item) && Array.isArray(t) ? e.item.replace("[", "").replace("]", "").split(",").map(s => s.trim()).forEach((s, a) => {
        i[s] = t[a]
    }) : /^\{.*\}$/.test(e.item) && !Array.isArray(t) && typeof t == "object" ? e.item.replace("{", "").replace("}", "").split(",").map(s => s.trim()).forEach(s => {
        i[s] = t[s]
    }) : i[e.item] = t, e.index && (i[e.index] = r), e.collection && (i[e.collection] = n), i
}

function ao(e) {
    return !Array.isArray(e) && !isNaN(e)
}

function On() {}
On.inline = (e, {
    expression: t
}, {
    cleanup: r
}) => {
    let n = X(e);
    n._x_refs || (n._x_refs = {}), n._x_refs[t] = e, r(() => delete n._x_refs[t])
};
d("ref", On);
d("if", (e, {
    expression: t
}, {
    effect: r,
    cleanup: n
}) => {
    e.tagName.toLowerCase() !== "template" && E("x-if can only be used on a <template> tag", e);
    let i = x(e, t),
        o = () => {
            if (e._x_currentIfEl) return e._x_currentIfEl;
            let a = e.content.cloneNode(!0).firstElementChild;
            return D(a, {}, e), m(() => {
                e.after(a), A(() => S(a))()
            }), e._x_currentIfEl = a, e._x_undoIf = () => {
                m(() => {
                    $(a), a.remove()
                }), delete e._x_currentIfEl
            }, a
        },
        s = () => {
            e._x_undoIf && (e._x_undoIf(), delete e._x_undoIf)
        };
    r(() => i(a => {
        a ? o() : s()
    })), n(() => e._x_undoIf && e._x_undoIf())
});
d("id", (e, {
    expression: t
}, {
    evaluate: r
}) => {
    r(t).forEach(i => gn(e, i))
});
H((e, t) => {
    e._x_ids && (t._x_ids = e._x_ids)
});
oe(Ie("@", $e(C("on:"))));
d("on", A((e, {
    value: t,
    modifiers: r,
    expression: n
}, {
    cleanup: i
}) => {
    let o = n ? x(e, n) : () => {};
    e.tagName.toLowerCase() === "template" && (e._x_forwardEvents || (e._x_forwardEvents = []), e._x_forwardEvents.includes(t) || e._x_forwardEvents.push(t));
    let s = le(e, t, r, a => {
        o(() => {}, {
            scope: {
                $event: a
            },
            params: [a]
        })
    });
    i(() => s())
}));
nt("Collapse", "collapse", "collapse");
nt("Intersect", "intersect", "intersect");
nt("Focus", "trap", "focus");
nt("Mask", "mask", "mask");

function nt(e, t, r) {
    d(t, n => E(`You can't use [x-${t}] without first installing the "${e}" plugin here: https://alpinejs.dev/plugins/${r}`, n))
}
K.setEvaluator(Jr);
K.setReactivityEngine({
    reactive: tt,
    effect: nn,
    release: on,
    raw: _
});
var qt = K;
window.Alpine = qt;
queueMicrotask(() => {
    qt.start()
});
})();
