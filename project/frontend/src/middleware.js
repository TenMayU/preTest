// middleware.js
import { NextResponse } from 'next/server'

const protectedRoutes = ['/', '/quote']

export function middleware(request) {
  const token = request.cookies.get('token')?.value
  const pathname = request.nextUrl.pathname

  const isProtected = protectedRoutes.some(
    (path) => pathname === path || pathname.startsWith(path + '/')
  )

  if (isProtected && !token) {
    const loginUrl = new URL('/login', request.url)
    return NextResponse.redirect(loginUrl)
  }

  return NextResponse.next()
}

export const config = {
  matcher: ['/', '/quote', '/quote/:path*'], // หรือใช้ ['/:path*'] ก็ได้
}
