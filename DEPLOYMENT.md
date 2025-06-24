# Vercel Deployment Guide for Go Backend

## Prerequisites

1. **Vercel Account**: Sign up at [vercel.com](https://vercel.com)
2. **Vercel CLI**: Install the Vercel CLI
   ```bash
   npm install -g vercel
   ```
3. **Git Repository**: Your code should be in a Git repository (GitHub, GitLab, or Bitbucket)

## Step-by-Step Deployment Guide

### Step 1: Prepare Your Environment Variables

Before deploying, you need to set up your database environment variables. You have several options:

#### Option A: Using Vercel Dashboard (Recommended)
1. Go to your Vercel dashboard
2. Create a new project
3. In the project settings, go to "Environment Variables"
4. Add the following variables:
   ```
   DB_HOST=your-database-host
   DB_USER=your-database-user
   DB_PASSWORD=your-database-password
   DB_NAME=your-database-name
   DB_PORT=3306
   ```

#### Option B: Using Vercel CLI
```bash
vercel env add DB_HOST
vercel env add DB_USER
vercel env add DB_PASSWORD
vercel env add DB_NAME
vercel env add DB_PORT
```

### Step 2: Deploy to Vercel

#### Method 1: Using Vercel CLI (Recommended)

1. **Navigate to your project directory**:
   ```bash
   cd model-api-backend
   ```

2. **Login to Vercel** (if not already logged in):
   ```bash
   vercel login
   ```

3. **Deploy your project**:
   ```bash
   vercel
   ```

4. **Follow the prompts**:
   - Link to existing project or create new
   - Confirm project name
   - Confirm deployment settings

5. **For production deployment**:
   ```bash
   vercel --prod
   ```

#### Method 2: Using Vercel Dashboard

1. **Push your code to Git repository** (GitHub, GitLab, or Bitbucket)

2. **Go to Vercel Dashboard**:
   - Visit [vercel.com/dashboard](https://vercel.com/dashboard)
   - Click "New Project"

3. **Import your repository**:
   - Select your Git provider
   - Choose your repository
   - Vercel will automatically detect it's a Go project

4. **Configure project settings**:
   - Framework Preset: Other
   - Build Command: `go build -o api/main api/main.go`
   - Output Directory: `.`
   - Install Command: `go mod download`

5. **Add environment variables** (as mentioned in Step 1)

6. **Deploy**:
   - Click "Deploy"

### Step 3: Verify Deployment

1. **Check deployment status** in Vercel dashboard
2. **Test your API endpoints**:
   ```bash
   curl https://your-project-name.vercel.app/api/health
   ```

3. **Check function logs** in Vercel dashboard if there are issues

### Step 4: Custom Domain (Optional)

1. **Go to your project settings** in Vercel dashboard
2. **Navigate to "Domains"**
3. **Add your custom domain**
4. **Configure DNS** as instructed by Vercel

## Important Notes

### Database Considerations

1. **Use a cloud database** (AWS RDS, Google Cloud SQL, PlanetScale, etc.)
2. **Ensure your database is accessible** from Vercel's servers
3. **Use connection pooling** for better performance
4. **Consider using a serverless database** like PlanetScale or Neon

### Environment Variables

Make sure to set these in your Vercel project:
- `DB_HOST`: Your database host
- `DB_USER`: Database username
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `DB_PORT`: Database port (usually 3306)

### Cold Starts

- Vercel functions have cold starts
- Consider using connection pooling
- Implement health checks for better monitoring

## Troubleshooting

### Common Issues

1. **Database Connection Errors**:
   - Check environment variables
   - Ensure database is accessible from Vercel
   - Check firewall settings

2. **Build Errors**:
   - Verify Go version compatibility
   - Check import paths
   - Ensure all dependencies are in go.mod

3. **Function Timeout**:
   - Vercel has a 10-second timeout for hobby plans
   - Consider upgrading to Pro plan for longer timeouts
   - Optimize database queries

### Debugging

1. **Check function logs** in Vercel dashboard
2. **Use Vercel CLI for local development**:
   ```bash
   vercel dev
   ```
3. **Test locally** before deploying:
   ```bash
   go run api/main.go
   ```

## Post-Deployment

1. **Monitor your application** using Vercel analytics
2. **Set up alerts** for errors
3. **Configure automatic deployments** from your main branch
4. **Set up preview deployments** for pull requests

## Support

- [Vercel Documentation](https://vercel.com/docs)
- [Vercel Go Runtime](https://vercel.com/docs/functions/serverless-functions/runtimes/go)
- [Vercel Community](https://github.com/vercel/vercel/discussions) 